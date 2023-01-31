package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
	"github.com/olahol/melody"
)

type TalkController struct{}

func RegisterTalkRouter(r *gin.RouterGroup) {
	talkController := &TalkController{}
	{
		r.POST("/get_talked_users", talkController.GetTalkedUsers)
	}

	// websocket
	{
		m := melody.New()
		r.GET("ws/:channel_id", func(c *gin.Context) {
			err := m.HandleRequest(c.Writer, c.Request)
			if err != nil {
				utils.ResponseError(c, error(common.UNKNOWNERROR))
			}
		})
		m.HandleMessage(func(s *melody.Session, message []byte) {
			req := &bo.SendMessageRequest{}
			err := json.Unmarshal(message, req)
			if err != nil {
				panic(err)
			}
			msg, err := service.GetTalkService().HandleMessage(s, req)
			if err != nil {
				panic(err)
			}
			err = m.BroadcastFilter(msg, func(q *melody.Session) bool {
				return q.Request.URL.Path == s.Request.URL.Path
			})
			if err != nil {
				panic(err)
			}
		})
	}
}

func (ins *TalkController) GetTalkedUsers(c *gin.Context) {
	req := &bo.GetTalkedUsersRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	resp, err := service.GetTalkService().GetTalkedUsers(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
