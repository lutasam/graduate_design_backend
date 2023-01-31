package dal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/repository"
	"sync"
)

type ActiveDal struct{}

var (
	activeDal     *ActiveDal
	activeDalOnce sync.Once
)

func GetActiveDal() *ActiveDal {
	activeDalOnce.Do(func() {
		activeDal = &ActiveDal{}
	})
	return activeDal
}

func (ins *ActiveDal) SetActiveCode(c *gin.Context, email, code string) error {
	_, err := repository.GetRedis().WithContext(c).Set(c, email+common.ACTIVECODESUFFIX, code, common.ACTIVECODEEXPTIME).Result()
	if err != nil {
		return common.REDISERROR
	}
	return nil
}

func (ins *ActiveDal) GetActiveCodeIfExist(c *gin.Context, email string) (string, error) {
	result, err := repository.GetRedis().WithContext(c).Get(c, email+common.ACTIVECODESUFFIX).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return "", common.USERINPUTERROR
	}
	if err != nil {
		return "", common.REDISERROR
	}
	return result, nil
}
