package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type CommentDal struct{}

var (
	commentDal     *CommentDal
	commentDalOnce sync.Once
)

func GetCommentDal() *CommentDal {
	commentDalOnce.Do(func() {
		commentDal = &CommentDal{}
	})
	return commentDal
}

func (ins *CommentDal) InsertComment(c *gin.Context, collectionName string, comment *model.Comment) (string, error) {
	res, err := repository.GetMongo().Database(model.Comment{}.DBName()).Collection(collectionName).InsertOne(c, comment)
	if err != nil {
		return "0", common.MONGOERROR
	}
	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (ins *CommentDal) FindComments(c *gin.Context, collectionName string) ([]*model.Comment, error) {
	var comments []*model.Comment
	cursor, err := repository.GetMongo().Database(model.Comment{}.DBName()).Collection(collectionName).Find(c, bson.D{},
		options.Find().SetSort(bson.D{{"created_at", -1}}))
	if err != nil {
		return nil, common.MONGOERROR
	}
	for cursor.Next(c) {
		comment := &model.Comment{}
		err := cursor.Decode(comment)
		if err != nil {
			return nil, common.MONGOERROR
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
