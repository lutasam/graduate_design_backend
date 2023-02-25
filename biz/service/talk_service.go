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
		talkedUserVOs = append(talkedUserVOs, &vo.TalkedUserVO{
			ID:            utils.Uint64ToString(user.ID),
			Avatar:        user.Avatar,
			Name:          user.Name,
			LastMessage:   lastMessage,
			PhoneNumber:   user.PhoneNumber,
			CreatedAt:     utils.TimeToString(lastMessageCreatedAt),
			MessageNumber: unreadMessageCount,
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
	if req.MessageType == common.NORMAL.Int() { // 只需返回一条message
		newMessage := &model.Message{
			UserID:     userID,
			UserName:   req.UserName,
			UserAvatar: req.UserAvatar,
			Content:    req.Content,
			CreatedAt:  time.Now(),
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
				CreatedAt:  utils.TimeToString(newMessage.CreatedAt),
			}},
		}, nil
	} else {
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
	}
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
			CreatedAt:  utils.TimeToString(msg.CreatedAt),
		})
	}
	return vos
}
