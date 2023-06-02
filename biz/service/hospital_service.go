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

type HospitalService struct{}

var (
	hospitalService     *HospitalService
	hospitalServiceOnce sync.Once
)

func GetHospitalService() *HospitalService {
	hospitalServiceOnce.Do(func() {
		hospitalService = &HospitalService{}
	})
	return hospitalService
}

func (ins *HospitalService) FindHospitals(c *gin.Context, req *bo.FindHospitalsRequest) (*bo.FindHospitalsResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 {
		return nil, common.USERINPUTERROR
	}
	hospitals, total, err := dal.GetHospitalDal().FindHospitals(c, req.CurrentPage, req.PageSize, req.HospitalName, req.City, req.HospitalRank, req.Characteristic)
	if err != nil {
		return nil, err
	}
	return &bo.FindHospitalsResponse{
		Total:     int(total),
		Hospitals: convertToHospitalVOs(hospitals),
	}, nil
}

func (ins *HospitalService) TakeHospitalInfo(c *gin.Context, req *bo.TakeHospitalInfoRequest) (*bo.TakeHospitalInfoResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	hospital, err := dal.GetHospitalDal().TakeHospitalByID(c, hospitalID)
	if err != nil {
		return nil, err
	}
	var rateScore float64
	if hospital.RatePeopleNum == 0 {
		rateScore = 0
	} else {
		rateScore = hospital.RateTotalScore / float64(hospital.RatePeopleNum)
	}
	return &bo.TakeHospitalInfoResponse{Hospital: &vo.HospitalVO{
		ID:             utils.Uint64ToString(hospital.ID),
		Name:           hospital.Name,
		City:           hospital.City,
		Address:        hospital.Address,
		HospitalRank:   common.ParseHospitalRank(hospital.HospitalRank).String(),
		Description:    hospital.Description,
		RateScore:      rateScore,
		RatePeopleNum:  hospital.RatePeopleNum,
		Characteristic: hospital.Characteristic,
		CreatedAt:      utils.TimeToDateString(hospital.CreatedAt),
	}}, nil
}

func (ins *HospitalService) FindHospitalDepartmentNames(c *gin.Context, req *bo.FindHospitalDepartmentNamesRequest) (*bo.FindHospitalDepartmentNamesResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	departments, err := dal.GetDepartmentDal().FindHospitalDepartments(c, hospitalID)
	if err != nil {
		return nil, err
	}
	return &bo.FindHospitalDepartmentNamesResponse{
		Total:           len(departments),
		DepartmentNames: convertToDepartmentVOs(departments),
	}, nil
}

func (ins *HospitalService) CreateHospital(c *gin.Context, req *bo.CreateHospitalRequest) (*bo.CreateHospitalResponse, error) {
	hospital := &model.Hospital{
		ID:             utils.GenerateHospitalID(),
		Name:           req.Name,
		City:           req.City,
		Address:        req.Address,
		HospitalRank:   req.HospitalRank,
		Description:    req.Description,
		Characteristic: req.Characteristic,
	}
	err := dal.GetHospitalDal().CreateHospital(c, hospital)
	if err != nil {
		return nil, err
	}
	for _, department := range req.Departments {
		department := &model.Department{
			ID:         utils.GenerateDepartmentID(),
			Name:       department.Name,
			Group:      department.Group,
			HospitalID: hospital.ID,
		}
		err := dal.GetDepartmentDal().CreateDepartment(c, department)
		if err != nil {
			return nil, err
		}
	}
	return &bo.CreateHospitalResponse{}, nil
}

func (ins *HospitalService) UpdateHospitalInfo(c *gin.Context, req *bo.UpdateHospitalInfoRequest) (*bo.UpdateHospitalInfoResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	hospital, err := dal.GetHospitalDal().TakeHospitalByID(c, hospitalID)
	if err != nil {
		return nil, err
	}
	if req.City != "" {
		hospital.City = req.City
	}
	if req.HospitalRank != 0 {
		hospital.HospitalRank = req.HospitalRank
	}
	if req.Name != "" {
		hospital.Name = req.Name
	}
	if req.Address != "" {
		hospital.Address = req.Address
	}
	if req.Description != "" {
		hospital.Description = req.Description
	}
	if req.Characteristic != "" {
		hospital.Characteristic = req.Characteristic
	}
	err = dal.GetHospitalDal().UpdateHospitalInfo(c, hospital)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateHospitalInfoResponse{}, nil
}

func (ins *HospitalService) DeleteHospital(c *gin.Context, req *bo.DeleteHospitalRequest) (*bo.DeleteHospitalResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetHospitalDal().TakeHospitalByID(c, hospitalID)
	if err != nil {
		return nil, err
	}
	// 考虑到外键，按顺序依次删除医院的departments -> 医院的doctors -> 医生对应的users -> hospital
	err = dal.GetDepartmentDal().DeleteDepartmentsByHospitalID(c, hospitalID)
	if err != nil {
		return nil, err
	}
	doctors, err := dal.GetDoctorDal().FindHospitalDoctors(c, hospitalID)
	if err != nil {
		return nil, err
	}
	var userIDs []uint64
	for _, doctor := range doctors {
		userIDs = append(userIDs, doctor.UserID)
	}
	err = dal.GetDoctorDal().DeleteDoctorsByUserIDs(c, userIDs)
	if err != nil {
		return nil, err
	}
	err = dal.GetUserDal().DeleteUsers(c, userIDs)
	if err != nil {
		return nil, err
	}
	err = dal.GetHospitalDal().DeleteHospital(c, hospitalID)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteHospitalResponse{}, nil
}

func (ins *HospitalService) UpdateHospitalRateScore(c *gin.Context, req *bo.UpdateHospitalRateScoreRequest) (*bo.UpdateHospitalRateScoreResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	isExist, err := dal.GetHospitalDal().IsUserRatedOnHospital(c, jwtStruct.UserID, hospitalID)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, common.USERHASRATED
	}
	_, err = dal.GetHospitalDal().TakeHospitalByID(c, hospitalID)
	if err != nil {
		return nil, err
	}
	err = dal.GetHospitalDal().UpdateHospitalRateScore(c, hospitalID, req.Score)
	if err != nil {
		return nil, err
	}
	err = dal.GetHospitalDal().SetUserRatedOnHospital(c, jwtStruct.UserID, hospitalID)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateHospitalRateScoreResponse{}, nil
}

func (ins *HospitalService) TakeHospitalRank(c *gin.Context, req *bo.TakeHospitalRankRequest) (*bo.TakeHospitalRankResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	rank, err := dal.GetHospitalDal().TakeHospitalRank(c, hospitalID, req.Area)
	if err != nil {
		return nil, err
	}
	return &bo.TakeHospitalRankResponse{Rank: rank}, nil
}

func (ins *HospitalService) SetHospitalReadCount(c *gin.Context, req *bo.SetHospitalReadCountRequest) (*bo.SetHospitalReadCountResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	err = dal.GetHospitalDal().SetHospitalReadCount(c, hospitalID)
	if err != nil {
		return nil, err
	}
	return &bo.SetHospitalReadCountResponse{}, nil
}

func (ins *HospitalService) TakeHospitalReadCount(c *gin.Context, req *bo.TakeHospitalReadCountRequest) (*bo.TakeHospitalReadCountResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	readCount, err := dal.GetHospitalDal().TakeHospitalReadCount(c, hospitalID)
	if err != nil {
		return nil, err
	}
	return &bo.TakeHospitalReadCountResponse{ReadCount: readCount}, nil
}

func convertToHospitalVOs(hospitals []*model.Hospital) []*vo.HospitalVO {
	var vos []*vo.HospitalVO
	for _, hospital := range hospitals {
		var rateScore float64
		if hospital.RatePeopleNum == 0 {
			rateScore = 0
		} else {
			rateScore = hospital.RateTotalScore / float64(hospital.RatePeopleNum)
		}
		vos = append(vos, &vo.HospitalVO{
			ID:             utils.Uint64ToString(hospital.ID),
			Name:           hospital.Name,
			City:           hospital.City,
			Address:        hospital.Address,
			HospitalRank:   common.ParseHospitalRank(hospital.HospitalRank).String(),
			Description:    hospital.Description,
			RateScore:      rateScore,
			RatePeopleNum:  hospital.RatePeopleNum,
			Characteristic: hospital.Characteristic,
			CreatedAt:      utils.TimeToDateString(hospital.CreatedAt),
		})
	}
	return vos
}

func convertToDepartmentVOs(departments []*model.Department) []*vo.DepartmentVO {
	var vos []*vo.DepartmentVO
	for _, department := range departments {
		vos = append(vos, &vo.DepartmentVO{
			Name:  department.Name,
			Group: department.Group,
		})
	}
	return vos
}
