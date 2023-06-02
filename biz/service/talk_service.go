package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/dal"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/utils"
	"github.com/lutasam/doctors/biz/vo"
	"github.com/olahol/melody"
	"strings"
	"sync"
	"time"
)

type TalkService struct{}

var (
	talkService     *TalkService
	talkServiceOnce sync.Once
)

func GetTalkService() *TalkService {
	talkServiceOnce.Do(func() {
		talkService = &TalkService{}
	})
	return talkService
}

func (ins *TalkService) GetTalkedUsers(c *gin.Context, req *bo.GetTalkedUsersRequest) (*bo.GetTalkedUsersResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	users, err := dal.GetTalkDal().FindTalkedUsers(c, jwtStruct.UserID)
	if err != nil {
		return nil, err
	}
	// 从mongo中获取对话数据
	var talkedUserVOs []*vo.TalkedUserVO
	for _, user := range users {
		minID, maxID := utils.MinUint64(jwtStruct.UserID, user.ID), utils.MaxUint64(jwtStruct.UserID, user.ID)
		minIDStr, maxIDStr := utils.Uint64ToString(minID), utils.Uint64ToString(maxID)
		collectionName := minIDStr + "_" + maxIDStr
		msgs, err := dal.GetTalkDal().FindMessages(collectionName)
		if err != nil {
			return nil, err
		}
		// 获取最新消息 从数组最后找对方说的消息
		var lastMessage string
		var lastMessageCreatedAt time.Time
		for i := len(msgs) - 1; i >= 0; i-- {
			if msgs[i].UserID != jwtStruct.UserID {
				lastMessage = msgs[i].Content
				lastMessageCreatedAt = msgs[i].CreatedAt
				break
			}
		}
		// 获取未读消息数量 从数组最后找，一旦已读标记则break循环
		unreadMessageCount := 0
		for i := len(msgs) - 1; i >= 0; i-- {
			if msgs[i].UserID != jwtStruct.UserID {
				if msgs[i].IsRead {
					break
				}
				unreadMessageCount++
			}
		}
		// 获取当前用户的在线状态
		isOnline, err := dal.GetTalkDal().GetUserOnlineStatus(c, user.ID)
		if err != nil {
			return nil, err
		}

		talkedUserVOs = append(talkedUserVOs, &vo.TalkedUserVO{
			ID:            utils.Uint64ToString(user.ID),
			Avatar:        user.Avatar,
			Name:          user.Name,
			LastMessage:   lastMessage,
			PhoneNumber:   user.PhoneNumber,
			CreatedAt:     utils.TimeToString(lastMessageCreatedAt),
			MessageNumber: unreadMessageCount,
			IsOnline:      isOnline,
		})
	}
	return &bo.GetTalkedUsersResponse{
		Total: len(talkedUserVOs),
		Users: talkedUserVOs,
	}, nil
}

func (ins *TalkService) AddTalkedUser(c *gin.Context, req *bo.AddTalkedUserRequest) (*bo.AddTalkedUserResponse, error) {
	toid, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	err = dal.GetTalkDal().AddTalkedUser(c, jwtStruct.UserID, toid)
	if err != nil {
		return nil, err
	}
	return &bo.AddTalkedUserResponse{}, nil
}

func (ins *TalkService) HandleMessage(s *melody.Session, req *bo.SendMessageRequest) (*bo.SendMessageResponse, error) {
	idx := strings.LastIndex(s.Request.URL.Path, "/")
	if idx == -1 {
		return nil, common.UNKNOWNERROR
	}
	collectionName := s.Request.URL.Path[idx+1:]
	userID, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	if req.MessageType == common.NORMAL.Int() { // 正常对话消息，只需返回一条message
		newMessage := &model.Message{
			UserID:     userID,
			UserName:   req.UserName,
			UserAvatar: req.UserAvatar,
			Content:    req.Content,
			CreatedAt:  time.Now(),
			ImageURL:   req.ImageURL,
			IsRead:     false,
		}
		err = dal.GetTalkDal().InsertMessage(collectionName, newMessage)
		if err != nil {
			return nil, err
		}
		return &bo.SendMessageResponse{
			Total: 1,
			Messages: []*vo.MessageVO{{
				ID:         newMessage.ID.String(),
				UserID:     utils.Uint64ToString(newMessage.UserID),
				UserName:   newMessage.UserName,
				UserAvatar: newMessage.UserAvatar,
				Content:    newMessage.Content,
				ImageURL:   req.ImageURL,
				CreatedAt:  utils.TimeToString(newMessage.CreatedAt),
			}},
		}, nil
	} else if req.MessageType == common.HISTORY.Int() { // 获取历史消息
		msgs, err := dal.GetTalkDal().FindMessages(collectionName)
		if err != nil {
			return nil, err
		}
		// 将对方的消息设置为已读
		var otherMessages []*model.Message
		for _, msg := range msgs {
			if msg.UserID != userID {
				otherMessages = append(otherMessages, msg)
			}
		}
		err = dal.GetTalkDal().UpdateMessagesStatusToRead(collectionName, otherMessages)
		if err != nil {
			return nil, err
		}
		return &bo.SendMessageResponse{
			Total:    len(msgs),
			Messages: convertToMessageVOs(msgs),
		}, nil
	} else { // 获取智能医生回答，AI的数据目前不往数据库存
		// 先从缓存中获取是否存在一样的问题，如果有直接返回缓存回答
		exist, answer, err := dal.GetTalkDal().TakeGPT2AnswerInCacheIfExist(req.Content)
		if err != nil {
			return nil, err
		}
		if exist { // 存在，则直接返回
			return &bo.SendMessageResponse{
				Total: 1,
				Messages: []*vo.MessageVO{
					{
						ID:         "-1",
						UserID:     "-1",
						UserName:   common.AI_DOCTOR_NAME,
						UserAvatar: common.AI_DOCTOR_AVATAR,
						Content:    answer, // 这里插入缓存的回答
						CreatedAt:  utils.TimeToString(time.Now()),
					},
				},
			}, nil
		}
		// 不存在，则通过python脚本输出回答
		answer, err = utils.GetGPT2Answer(req.Content)
		if err != nil {
			return nil, err
		}
		// 当问题不为空时，再把该问题和回答送入缓存中(防止用户传输图片等非文字类内容)
		if req.Content != "" {
			err = dal.GetTalkDal().AddGPT2AnswerToCache(req.Content, answer)
			if err != nil {
				return nil, err
			}
		}
		return &bo.SendMessageResponse{
			Total: 1,
			Messages: []*vo.MessageVO{
				{
					ID:         "-1", // 消息Mongo中的id，由于不存在数据库所以没有这个ID
					UserID:     "-1",
					UserName:   common.AI_DOCTOR_NAME,
					UserAvatar: common.AI_DOCTOR_AVATAR,
					Content:    answer, // 这里插入机器人的回答
					CreatedAt:  utils.TimeToString(time.Now()),
				},
			},
		}, nil
	}
}

func (ins *TalkService) SetUserOnline(c *gin.Context, req *bo.SetUserOnlineRequest) (*bo.SetUserOnlineResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	err = dal.GetTalkDal().SetUserOnline(c, jwtStruct.UserID)
	if err != nil {
		return nil, err
	}
	return &bo.SetUserOnlineResponse{}, nil
}

func (ins *TalkService) SetUserOffline(c *gin.Context, req *bo.SetUserOfflineRequest) (*bo.SetUserOfflineResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	err = dal.GetTalkDal().SetUserOffline(c, jwtStruct.UserID)
	if err != nil {
		return nil, err
	}
	return &bo.SetUserOfflineResponse{}, nil
}

func convertToMessageVOs(messages []*model.Message) []*vo.MessageVO {
	var vos []*vo.MessageVO
	for _, msg := range messages {
		vos = append(vos, &vo.MessageVO{
			ID:         msg.ID.String(),
			UserID:     utils.Uint64ToString(msg.UserID),
			UserName:   msg.UserName,
			UserAvatar: msg.UserAvatar,
			Content:    msg.Content,
			ImageURL:   msg.ImageURL,
			CreatedAt:  utils.TimeToString(msg.CreatedAt),
		})
	}
	return vos
}
