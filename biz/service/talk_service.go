package service

import (
	"encoding/json"
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
	talkedUserVOs := convert2TalkedUserVOs(users)
	return &bo.GetTalkedUsersResponse{
		Total: len(talkedUserVOs),
		Users: talkedUserVOs,
	}, nil
}

func (ins *TalkService) HandleMessage(s *melody.Session, req *bo.SendMessageRequest) ([]byte, error) {
	var bytes []byte
	idx := strings.LastIndex(s.Request.URL.Path, "/")
	if idx == -1 {
		return nil, common.UNKNOWNERROR
	}
	collectionName := s.Request.URL.Path[idx+1:]
	if req.MessageType == common.NORMAL.Int() { // 只需返回一条message
		userID, err := utils.StringToUint64(req.UserID)
		if err != nil {
			return nil, err
		}
		err = dal.GetTalkDal().InsertMessage(collectionName, &model.Message{
			UserID:     userID,
			UserName:   req.UserName,
			UserAvatar: req.UserAvatar,
			Content:    req.Content,
			CreatedAt:  time.Now(),
		})
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(req)
		if err != nil {
			return nil, common.USERINPUTERROR
		}
	} else {
		msgs, err := dal.GetTalkDal().FindMessages(collectionName)
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(msgs)
		if err != nil {
			return nil, common.UNKNOWNERROR
		}
	}
	return bytes, nil
}

func convert2TalkedUserVOs(users []*model.User) []*vo.TalkedUser {
	var vos []*vo.TalkedUser
	for _, user := range users {
		vos = append(vos, &vo.TalkedUser{
			ID:     utils.Uint64ToString(user.ID),
			Name:   user.Name,
			Avatar: user.Avatar,
		})
	}
	return vos
}
