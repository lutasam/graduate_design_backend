package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/dal"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/utils"
	"github.com/lutasam/doctors/biz/vo"
	"sync"
)

type UserService struct{}

var (
	userService     *UserService
	userServiceOnce sync.Once
)

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (ins *UserService) TakeUserInfo(c *gin.Context, req *bo.TakeUserInfoRequest) (*bo.TakeUserInfoResponse, error) {
	var userID uint64
	if req.UserID == "" {
		jwtStruct, err := utils.GetCtxUserInfoJWT(c)
		if err != nil {
			return nil, err
		}
		userID = jwtStruct.UserID
	} else {
		id, err := utils.StringToUint64(req.UserID)
		if err != nil {
			return nil, err
		}
		userID = id
	}
	user, err := dal.GetUserDal().TakeUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	return &bo.TakeUserInfoResponse{User: &vo.UserVO{
		ID:            utils.Uint64ToString(user.ID),
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		Name:          user.Name,
		Avatar:        user.Avatar,
		Birthday:      utils.TimeToDateString(user.Birthday),
		CharacterType: common.ParseCharacterType(user.CharacterType).String(),
	}}, nil
}

func (ins *UserService) UpdateUserInfo(c *gin.Context, req *bo.UpdateUserInfoRequest) (*bo.UpdateUserInfoResponse, error) {
	if req.PhoneNumber != "" && !utils.IsValidPhoneNumber(req.PhoneNumber) ||
		req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return nil, common.USERINPUTERROR
	}
	birthday, err := utils.DateStringToTime(req.Birthday)
	if req.Birthday != "" && err != nil {
		return nil, err
	}
	var userID uint64
	if req.UserID == "" {
		jwtStruct, err := utils.GetCtxUserInfoJWT(c)
		if err != nil {
			return nil, err
		}
		userID = jwtStruct.UserID
	} else {
		id, err := utils.StringToUint64(req.UserID)
		if err != nil {
			return nil, err
		}
		userID = id
	}
	user, err := dal.GetUserDal().TakeUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Birthday != "" {
		user.Birthday = birthday
	}
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateUserInfoResponse{}, nil
}

func (ins *UserService) FindUsers(c *gin.Context, req *bo.FindUsersRequest) (*bo.FindUsersResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 {
		return nil, common.USERINPUTERROR
	}
	users, err := dal.GetUserDal().FindUsers(c, req.CurrentPage, req.PageSize, req.Name)
	if err != nil {
		return nil, err
	}
	return &bo.FindUsersResponse{
		Total: len(users),
		Users: convertToUserVOs(users),
	}, nil
}

func (ins *UserService) DeleteUser(c *gin.Context, req *bo.DeleteUserRequest) (*bo.DeleteUserResponse, error) {
	id, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetUserDal().TakeUserByID(c, id)
	if err != nil {
		return nil, err
	}
	err = dal.GetUserDal().DeleteUser(c, id)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteUserResponse{}, nil
}

func convertToUserVOs(users []*model.User) []*vo.UserVO {
	var userVOs []*vo.UserVO
	for _, user := range users {
		userVOs = append(userVOs, &vo.UserVO{
			ID:            utils.Uint64ToString(user.ID),
			Email:         user.Email,
			PhoneNumber:   user.PhoneNumber,
			Name:          user.Name,
			Birthday:      utils.TimeToDateString(user.Birthday),
			Avatar:        user.Avatar,
			CharacterType: common.ParseCharacterType(user.CharacterType).String(),
		})
	}
	return userVOs
}
