package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/dal"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/utils"
	"github.com/lutasam/doctors/biz/vo"
	"sync"
	"time"
)

type CommentService struct{}

var (
	commentService     *CommentService
	commentServiceOnce sync.Once
)

func GetCommentService() *CommentService {
	commentServiceOnce.Do(func() {
		commentService = &CommentService{}
	})
	return commentService
}

func (ins *CommentService) AddComment(c *gin.Context, req *bo.AddCommentRequest) (*bo.AddCommentResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByID(c, jwtStruct.UserID)
	if err != nil {
		return nil, err
	}
	comment := &model.Comment{
		UserName:  user.Name,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	id, err := dal.GetCommentDal().InsertComment(c, req.CommentAreaName, comment)
	if err != nil {
		return nil, err
	}
	return &bo.AddCommentResponse{Comment: &vo.CommentVO{
		ID:        id,
		UserName:  comment.UserName,
		Content:   comment.Content,
		CreatedAt: utils.TimeToString(comment.CreatedAt),
	}}, nil
}

func (ins *CommentService) FindComments(c *gin.Context, req *bo.FindCommentsRequest) (*bo.FindCommentsResponse, error) {
	comments, err := dal.GetCommentDal().FindComments(c, req.CommentAreaName)
	if err != nil {
		return nil, err
	}
	return &bo.FindCommentsResponse{
		Total:    len(comments),
		Comments: convertToCommentVOs(comments),
	}, nil
}

func convertToCommentVOs(comments []*model.Comment) []*vo.CommentVO {
	var vos []*vo.CommentVO
	for _, comment := range comments {
		vos = append(vos, &vo.CommentVO{
			ID:        comment.ID.String(),
			UserName:  comment.UserName,
			Content:   comment.Content,
			CreatedAt: utils.TimeToString(comment.CreatedAt),
		})
	}
	return vos
}
