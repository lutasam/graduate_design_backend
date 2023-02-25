package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/middleware"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
	"github.com/olahol/melody"
	"net/http"
)

type TalkController struct{}

func RegisterTalkRouter(r *gin.RouterGroup) {
	talkController := &TalkController{}
	{
		r.POST("/get_talked_users", middleware.JWTAuth(), talkController.GetTalkedUsers)
		r.POST("/add_talked_user", middleware.JWTAuth(), talkController.AddTalkedUser)
	}

	// websocket
	{
		m := melody.New()
		m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true } // 跨域
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
			resp, err := service.GetTalkService().HandleMessage(s, req)
			if err != nil {
				panic(err)
			}
			bytes, err := json.Marshal(resp)
			if err != nil {
				panic(err)
			}
			err = m.BroadcastFilter(bytes, func(q *melody.Session) bool {
				return q.Request.URL.Path == s.Request.URL.Path
			})
			if err != nil {
				panic(err)
			}
		})
	}
}

func (ins *TalkController) AddTalkedUser(c *gin.Context) {
	req := &bo.AddTalkedUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	resp, err := service.GetTalkService().AddTalkedUser(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
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
