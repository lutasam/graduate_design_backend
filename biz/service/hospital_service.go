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
	hospitals, total, err := dal.GetHospitalDal().FindHospitals(c, req.CurrentPage, req.PageSize, req.HospitalName, req.City, req.HospitalRank)
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
	return &bo.TakeHospitalInfoResponse{Hospital: &vo.HospitalVO{
		ID:           utils.Uint64ToString(hospital.ID),
		Name:         hospital.Name,
		City:         hospital.City,
		Address:      hospital.Address,
		HospitalRank: common.ParseHospitalRank(hospital.HospitalRank).String(),
		Description:  hospital.Description,
		CreatedAt:    utils.TimeToDateString(hospital.CreatedAt),
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
	var names []string
	for _, department := range departments {
		names = append(names, department.Name)
	}
	return &bo.FindHospitalDepartmentNamesResponse{
		Total:           len(names),
		DepartmentNames: names,
	}, nil
}

func convertToHospitalVOs(hospitals []*model.Hospital) []*vo.HospitalVO {
	var vos []*vo.HospitalVO
	for _, hospital := range hospitals {
		vos = append(vos, &vo.HospitalVO{
			ID:           utils.Uint64ToString(hospital.ID),
			Name:         hospital.Name,
			City:         hospital.City,
			Address:      hospital.Address,
			HospitalRank: common.ParseHospitalRank(hospital.HospitalRank).String(),
			Description:  hospital.Description,
			CreatedAt:    utils.TimeToDateString(hospital.CreatedAt),
		})
	}
	return vos
}
