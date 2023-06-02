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

type DoctorService struct{}

var (
	doctorService     *DoctorService
	doctorServiceOnce sync.Once
)

func GetDoctorService() *DoctorService {
	doctorServiceOnce.Do(func() {
		doctorService = &DoctorService{}
	})
	return doctorService
}

func (ins *DoctorService) TakeDoctorInfo(c *gin.Context, req *bo.TakeDoctorInfoRequest) (*bo.TakeDoctorInfoResponse, error) {
	var doctor *model.Doctor
	if req.DoctorID == "" {
		jwtStruct, err := utils.GetCtxUserInfoJWT(c)
		if err != nil {
			return nil, err
		}
		doctor, err = dal.GetDoctorDal().TakeDoctorByUserID(c, jwtStruct.UserID)
		if err != nil {
			return nil, err
		}
	} else {
		doctorID, err := utils.StringToUint64(req.DoctorID)
		if err != nil {
			return nil, err
		}
		doctor, err = dal.GetDoctorDal().TakeDoctorByID(c, doctorID)
		if err != nil {
			return nil, err
		}
	}
	hospital, err := dal.GetHospitalDal().TakeHospitalByID(c, doctor.HospitalID)
	if err != nil {
		return nil, err
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByID(c, doctor.DepartmentID)
	if err != nil {
		return nil, err
	}
	// 评分注意评分人数未0时不能直接除
	var rateScore float64
	if doctor.RatePeopleNum == 0 {
		rateScore = 0
	} else {
		rateScore = doctor.RateTotalScore / float64(doctor.RatePeopleNum)
	}
	return &bo.TakeDoctorInfoResponse{Doctor: &vo.DoctorVO{
		ID:               utils.Uint64ToString(doctor.ID),
		Name:             doctor.User.Name,
		HospitalName:     hospital.Name,
		DepartmentName:   department.Name,
		Avatar:           doctor.User.Avatar,
		ProfessionalRank: common.ParseProfessionalRank(doctor.ProfessionalRank).String(),
		StudyDirection:   doctor.StudyDirection,
		Description:      doctor.Description,
		RateScore:        rateScore,
		RatePeopleNum:    doctor.RatePeopleNum,
	}}, nil
}

func (ins *DoctorService) UpdateDoctorInfo(c *gin.Context, req *bo.UpdateDoctorInfoRequest) (*bo.UpdateDoctorInfoResponse, error) {
	var doctor *model.Doctor
	if req.DoctorID == "" {
		jwtStruct, err := utils.GetCtxUserInfoJWT(c)
		if err != nil {
			return nil, err
		}
		doctor, err = dal.GetDoctorDal().TakeDoctorByUserID(c, jwtStruct.UserID)
		if err != nil {
			return nil, err
		}
	} else {
		id, err := utils.StringToUint64(req.DoctorID)
		if err != nil {
			return nil, err
		}
		doctor, err = dal.GetDoctorDal().TakeDoctorByID(c, id)
		if err != nil {
			return nil, err
		}
	}
	if req.HospitalName != "" {
		hospital, err := dal.GetHospitalDal().TakeHospitalByName(c, req.HospitalName)
		if err != nil {
			return nil, err
		}
		doctor.HospitalID = hospital.ID
	}
	if req.StudyDirection != "" {
		doctor.StudyDirection = req.StudyDirection
	}
	if req.Description != "" {
		doctor.Description = req.Description
	}
	if req.ProfessionalRank != 0 {
		doctor.ProfessionalRank = req.ProfessionalRank
	}
	if req.DepartmentName != "" {
		department, err := dal.GetDepartmentDal().TakeDepartmentByName(c, req.DepartmentName)
		if err != nil {
			return nil, err
		}
		doctor.DepartmentID = department.ID
	}
	err := dal.GetDoctorDal().UpdateDoctor(c, doctor)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateDoctorInfoResponse{}, nil
}

func (ins *DoctorService) FindDoctors(c *gin.Context, req *bo.FindDoctorsRequest) (*bo.FindDoctorsResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 {
		return nil, common.USERINPUTERROR
	}
	doctors, total, err := dal.GetDoctorDal().FindDoctors(c, req.CurrentPage, req.PageSize, req.StudyDirection, req.HospitalName, req.ProfessionalRank, req.Department)
	if err != nil {
		return nil, err
	}
	return &bo.FindDoctorsResponse{
		Total:   int(total),
		Doctors: convertToDoctorVOs(doctors),
	}, nil
}

func (ins *DoctorService) DeleteDoctor(c *gin.Context, req *bo.DeleteDoctorRequest) (*bo.DeleteDoctorResponse, error) {
	doctorID, err := utils.StringToUint64(req.DoctorID)
	if err != nil {
		return nil, err
	}
	doctor, err := dal.GetDoctorDal().TakeDoctorByID(c, doctorID)
	if err != nil {
		return nil, err
	}
	err = dal.GetDoctorDal().DeleteDoctor(c, doctorID)
	if err != nil {
		return nil, err
	}
	err = dal.GetUserDal().DeleteUser(c, doctor.UserID)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteDoctorResponse{}, nil
}

func (ins *DoctorService) ActiveDoctor(c *gin.Context, req *bo.ActiveDoctorRequest) (*bo.ActiveDoctorResponse, error) {
	doctorID, err := utils.StringToUint64(req.DoctorID)
	if err != nil {
		return nil, err
	}
	err = dal.GetDoctorDal().ActiveDoctor(c, doctorID)
	if err != nil {
		return nil, err
	}
	return &bo.ActiveDoctorResponse{}, nil
}

func (ins *DoctorService) FindHospitalDoctors(c *gin.Context, req *bo.FindHospitalDoctorsRequest) (*bo.FindHospitalDoctorsResponse, error) {
	hospitalID, err := utils.StringToUint64(req.HospitalID)
	if err != nil {
		return nil, err
	}
	doctors, err := dal.GetDoctorDal().FindHospitalDoctors(c, hospitalID)
	if err != nil {
		return nil, err
	}
	return &bo.FindHospitalDoctorsResponse{
		Total:   len(doctors),
		Doctors: convertToDoctorVOs(doctors),
	}, nil
}

func (ins *DoctorService) TakeDoctorRank(c *gin.Context, req *bo.TakeDoctorRankRequest) (*bo.TakeDoctorRankResponse, error) {
	doctorID, err := utils.StringToUint64(req.DoctorID)
	if err != nil {
		return nil, err
	}
	rank, err := dal.GetDoctorDal().TakeDoctorRank(c, doctorID, req.DepartmentName)
	if err != nil {
		return nil, err
	}
	return &bo.TakeDoctorRankResponse{Rank: rank}, nil
}

func (ins *DoctorService) SetDoctorReadCount(c *gin.Context, req *bo.SetDoctorReadCountRequest) (*bo.SetDoctorReadCountResponse, error) {
	doctorID, err := utils.StringToUint64(req.DoctorID)
	if err != nil {
		return nil, err
	}
	err = dal.GetDoctorDal().SetDoctorReadCount(c, doctorID)
	if err != nil {
		return nil, err
	}
	return &bo.SetDoctorReadCountResponse{}, nil
}

func (ins *DoctorService) TakeDoctorReadCount(c *gin.Context, req *bo.TakeDoctorReadCountRequest) (*bo.TakeDoctorReadCountResponse, error) {
	doctorID, err := utils.StringToUint64(req.DoctorID)
	if err != nil {
		return nil, err
	}
	readCount, err := dal.GetDoctorDal().TakeDoctorReadCount(c, doctorID)
	if err != nil {
		return nil, err
	}
	return &bo.TakeDoctorReadCountResponse{ReadCount: readCount}, nil
}

func (ins *DoctorService) UpdateDoctorRateScore(c *gin.Context, req *bo.UpdateDoctorRateScoreRequest) (*bo.UpdateDoctorRateScoreResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	doctorID, err := utils.StringToUint64(req.DoctorID)
	if err != nil {
		return nil, err
	}
	isExist, err := dal.GetDoctorDal().IsUserRatedOnDoctor(c, jwtStruct.UserID, doctorID)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, common.USERHASRATED
	}
	_, err = dal.GetDoctorDal().TakeDoctorByID(c, doctorID)
	if err != nil {
		return nil, err
	}
	err = dal.GetDoctorDal().UpdateDoctorRateScore(c, doctorID, req.Score)
	if err != nil {
		return nil, err
	}
	err = dal.GetDoctorDal().SetUserRatedOnDoctor(c, jwtStruct.UserID, doctorID)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateDoctorRateScoreResponse{}, nil
}

func convertToDoctorVOs(doctors []*model.Doctor) []*vo.DoctorVO {
	var doctorVOs []*vo.DoctorVO
	for _, doctor := range doctors {
		var rateScore float64
		if doctor.RatePeopleNum == 0 {
			rateScore = 0
		} else {
			rateScore = doctor.RateTotalScore / float64(doctor.RatePeopleNum)
		}
		doctorVOs = append(doctorVOs, &vo.DoctorVO{
			ID:               utils.Uint64ToString(doctor.ID),
			Name:             doctor.User.Name,
			HospitalName:     doctor.Hospital.Name,
			DepartmentName:   doctor.Department.Name,
			Avatar:           doctor.User.Avatar,
			ProfessionalRank: common.ParseProfessionalRank(doctor.ProfessionalRank).String(),
			StudyDirection:   doctor.StudyDirection,
			Description:      doctor.Description,
			RateScore:        rateScore,
			RatePeopleNum:    doctor.RatePeopleNum,
		})
	}
	return doctorVOs
}
