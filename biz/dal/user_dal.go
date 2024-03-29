package dal

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"github.com/lutasam/doctors/biz/utils"
	"sync"
)

type UserDal struct{}

var (
	userDal     *UserDal
	userDalOnce sync.Once
)

func GetUserDal() *UserDal {
	userDalOnce.Do(func() {
		userDal = &UserDal{}
	})
	return userDal
}

func (ins *UserDal) TakeUserByID(c *gin.Context, userID uint64) (*model.User, error) {
	user := &model.User{}
	userJSON, err := repository.GetRedis().Get(c, utils.Uint64ToString(userID)+common.USERINFOIDSUFFIX).Result()
	if err != nil && errors.Is(err, redis.Nil) || user.ID == 0 { // redis未命中
		err := repository.GetDB().WithContext(c).Table(user.TableName()).Where("id = ?", userID).Find(user).Error
		if err != nil {
			return nil, common.DATABASEERROR
		}
		if user.ID == 0 {
			return nil, common.USERDOESNOTEXIST
		}
		go func() {
			j, err := json.Marshal(user)
			if err != nil {
				panic(err)
			}
			err = repository.GetRedis().Set(c, utils.Uint64ToString(userID)+common.USERINFOIDSUFFIX, j, common.REDISEXPIRETIME).Err()
			if err != nil {
				panic(err)
			}
		}()
		return user, nil
	}
	if err != nil {
		return nil, common.REDISERROR
	}
	err = json.Unmarshal([]byte(userJSON), user)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return user, nil
	//user := &model.User{}
	//err := repository.GetDB().WithContext(c).Table(user.TableName()).Where("id = ?", userID).Find(user).Error
	//if err != nil {
	//	return nil, common.DATABASEERROR
	//}
	//if user.ID == 0 {
	//	return nil, common.USERDOESNOTEXIST
	//}
	//return user, nil
}

func (ins *UserDal) TakeUserByEmail(c *gin.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) FindUsersByIDs(c *gin.Context, userIDs []uint64) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).Where("id in ?", userIDs).Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) FindUsers(c *gin.Context, currentPage, pageSize int, name string) ([]*model.User, error) {
	var users []*model.User
	sql := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).Where("id != ?", 0)
	if name != "" {
		sql = sql.Where("name like ?", "%"+name+"%")
	}
	err := sql.Offset(pageSize * (currentPage - 1)).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) CreateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Create(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) UpdateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Updates(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	go func() {
		err = repository.GetRedis().Del(c, utils.Uint64ToString(user.ID)+common.USERINFOIDSUFFIX).Err()
		if err != nil && !errors.Is(err, redis.Nil) {
			panic(err)
		}
	}()
	return nil
}

func (ins *UserDal) DeleteUser(c *gin.Context, userID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).Where("id = ?", userID).Delete(&model.User{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	go func() {
		err = repository.GetRedis().Del(c, utils.Uint64ToString(userID)+common.USERINFOIDSUFFIX).Err()
		if err != nil && !errors.Is(err, redis.Nil) {
			panic(err)
		}
	}()
	return nil
}

func (ins *UserDal) DeleteUsers(c *gin.Context, userIDs []uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).Where("id in ?", userIDs).Delete(&model.User{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
