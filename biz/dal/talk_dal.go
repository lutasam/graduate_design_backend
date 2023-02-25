package dal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"github.com/lutasam/doctors/biz/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type TalkDal struct{}

var (
	talkDal     *TalkDal
	talkDalOnce sync.Once
)

func GetTalkDal() *TalkDal {
	talkDalOnce.Do(func() {
		talkDal = &TalkDal{}
	})
	return talkDal
}

func (ins *TalkDal) AddTalkedUser(c *gin.Context, fromUserID, toUserID uint64) error {
	fromid, toid := utils.Uint64ToString(fromUserID), utils.Uint64ToString(toUserID)
	_, err := repository.GetRedis().WithContext(c).SAdd(c, fromid+common.TALKEDUSERLISTSUFFIX, toid).Result()
	if err != nil {
		return common.REDISERROR
	}
	_, err = repository.GetRedis().WithContext(c).SAdd(c, toid+common.TALKEDUSERLISTSUFFIX, fromid).Result()
	if err != nil {
		return common.REDISERROR
	}
	return nil
}

func (ins *TalkDal) FindTalkedUsers(c *gin.Context, userID uint64) ([]*model.User, error) {
	id := utils.Uint64ToString(userID)
	userIDs, err := repository.GetRedis().WithContext(c).SMembers(c, id+common.TALKEDUSERLISTSUFFIX).Result()
	if err != nil && err != redis.Nil {
		return nil, common.REDISERROR
	}
	var ids []uint64
	for i := range userIDs {
		id, err := utils.StringToUint64(userIDs[i])
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	users, err := GetUserDal().FindUsersByIDs(c, ids)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ins *TalkDal) InsertMessage(collectionName string, msg *model.Message) error {
	c := context.Background()
	_, err := repository.GetMongo().Database(model.Message{}.DBName()).Collection(collectionName).InsertOne(c, msg)
	if err != nil {
		return common.MONGOERROR
	}
	return nil
}

func (ins *TalkDal) FindMessages(collectionName string) ([]*model.Message, error) {
	var msgs []*model.Message
	c := context.Background()
	cursor, err := repository.GetMongo().Database(model.Message{}.DBName()).Collection(collectionName).Find(c, bson.D{},
		options.Find().SetSort(bson.D{{"created_at", 1}}))
	if err != nil {
		return nil, common.MONGOERROR
	}
	for cursor.Next(c) {
		msg := &model.Message{}
		err := cursor.Decode(msg)
		if err != nil {
			return nil, common.MONGOERROR
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

func (ins *TalkDal) UpdateMessagesStatusToRead(collectionName string, msgs []*model.Message) error {
	c := context.Background()
	for _, msg := range msgs {
		_, err := repository.GetMongo().Database(model.Message{}.DBName()).Collection(collectionName).
			UpdateByID(c, msg.ID, bson.D{{"$set", bson.D{{"is_read", true}}}})
		if err != nil {
			return common.MONGOERROR
		}
	}
	return nil
}
