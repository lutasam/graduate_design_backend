package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/dal"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/utils"
	"sync"
	"time"
)

type LoginService struct{}

var (
	loginService     *LoginService
	loginServiceOnce sync.Once
)

func GetLoginService() *LoginService {
	loginServiceOnce.Do(func() {
		loginService = &LoginService{}
	})
	return loginService
}

func (ins *LoginService) Login(c *gin.Context, req *bo.LoginRequest) (*bo.LoginResponse, error) {
	if !utils.IsValidEmail(req.Email) || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	err = utils.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}
	jwt, err := utils.GenerateJWTByUserInfo(user)
	if err != nil {
		return nil, err
	}
	return &bo.LoginResponse{
		UserID:        utils.Uint64ToString(user.ID),
		CharacterType: user.CharacterType,
		Token:         jwt,
	}, nil
}

func (ins *LoginService) ApplyRegister(c *gin.Context, req *bo.ApplyRegisterRequest) (*bo.ApplyRegisterResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	_, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil && errors.Is(err, common.DATABASEERROR) {
		return nil, err
	}
	if err == nil { // username is duplicate with other guys
		return nil, common.USEREXISTED
	}
	go func() {
		err := sendActiveUserEmail(c, req.Email)
		if err != nil {
			panic(err)
		}
	}()
	return &bo.ApplyRegisterResponse{}, nil
}

func (ins *LoginService) ActiveUser(c *gin.Context, req *bo.ActiveUserRequest) (*bo.ActiveUserResponse, error) {
	if !utils.IsValidEmail(req.Email) || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	code, err := dal.GetActiveDal().GetActiveCodeIfExist(c, req.Email)
	if err != nil {
		return nil, err
	}
	if code != req.ActiveCode {
		return nil, common.ACTIVECODEERROR
	}
	encryptPass, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	id := utils.GenerateUserID()
	namePrefix := common.ParseCharacterType(req.CharacterType).String()
	user := &model.User{
		ID:            id,
		Email:         req.Email,
		PhoneNumber:   "",
		Password:      encryptPass,
		Name:          namePrefix + utils.Uint64ToString(id),
		Birthday:      time.Now(),
		Avatar:        common.DEFAULTAVATARURL,
		CharacterType: req.CharacterType,
		Sex:           common.MALE.Int(),
		City:          "",
		Address:       "",
	}
	err = dal.GetUserDal().CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	if req.CharacterType == common.DOCTOR.Int() {
		doctorID := utils.GenerateDoctorID()
		err = dal.GetDoctorDal().CreateDoctor(c, &model.Doctor{
			ID:       doctorID,
			UserID:   user.ID,
			IsActive: false,
		})
		if err != nil {
			return nil, err
		}
	}
	return &bo.ActiveUserResponse{}, nil
}

func (ins *LoginService) ResetPassword(c *gin.Context, req *bo.ResetPasswordRequest) (*bo.ResetPasswordResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	_, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	go func() {
		err := sendActiveUserEmail(c, req.Email)
		if err != nil {
			return
		}
	}()
	return &bo.ResetPasswordResponse{}, nil
}

func (ins *LoginService) ActiveResetPassword(c *gin.Context, req *bo.ActiveResetPasswordRequest) (*bo.ActiveResetPasswordResponse, error) {
	if !utils.IsValidEmail(req.Email) || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	code, err := dal.GetActiveDal().GetActiveCodeIfExist(c, req.Email)
	if err != nil {
		return nil, err
	}
	if code != req.ActiveCode {
		return nil, common.ACTIVECODEERROR
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	encryptPass, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encryptPass
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.ActiveResetPasswordResponse{}, nil
}

func (ins *LoginService) ApplyChangeUserEmail(c *gin.Context, req *bo.ApplyChangeUserEmailRequest) (*bo.ApplyChangeUserEmailResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	_, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil && errors.Is(err, common.DATABASEERROR) {
		return nil, err
	}
	if err == nil {
		return nil, common.USEREXISTED
	}
	go func() {
		err := sendActiveUserEmail(c, req.Email)
		if err != nil {
			panic(err)
		}
	}()
	return &bo.ApplyChangeUserEmailResponse{}, nil
}

func (ins *LoginService) ActiveChangeUserEmail(c *gin.Context, req *bo.ActiveChangeUserEmailRequest) (*bo.ActiveChangeUserEmailResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	code, err := dal.GetActiveDal().GetActiveCodeIfExist(c, req.Email)
	if err != nil {
		return nil, err
	}
	if code != req.ActiveCode {
		return nil, common.ACTIVECODEERROR
	}
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, userInfo.Email)
	if err != nil {
		return nil, err
	}
	user.Email = req.Email
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.ActiveChangeUserEmailResponse{}, nil
}

func sendActiveUserEmail(c *gin.Context, email string) error {
	activeCode := utils.GenerateActiveCode()
	err := dal.GetActiveDal().SetActiveCode(c, email, activeCode)
	if err != nil {
		return err
	}
	subject := "[验证激活码]找大夫在线咨询网"
	body := `
验证码：%s。5分钟之内有效。<br>
如果不是您本人操作，请忽视该邮件。
`
	err = utils.SendMail(email, subject, fmt.Sprintf(body, activeCode))
	if err != nil {
		return err
	}
	return nil
}
