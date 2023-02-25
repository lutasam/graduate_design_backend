package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"gorm.io/gorm/clause"
	"sync"
)

type DoctorDal struct{}

var (
	doctorDal     *DoctorDal
	doctorDalOnce sync.Once
)

func GetDoctorDal() *DoctorDal {
	doctorDalOnce.Do(func() {
		doctorDal = &DoctorDal{}
	})
	return doctorDal
}

func (ins *DoctorDal) CreateDoctor(c *gin.Context, doctor *model.Doctor) error {
	err := repository.GetDB().WithContext(c).Table(doctor.TableName()).Create(doctor).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) TakeDoctorByID(c *gin.Context, doctorID uint64) (*model.Doctor, error) {
	doctor := &model.Doctor{}
	err := repository.GetDB().WithContext(c).Model(&model.Doctor{}).Preload(clause.Associations).
		Where("doctors.id = ?", doctorID).Find(doctor).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if doctor.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return doctor, nil
}

func (ins *DoctorDal) TakeDoctorByUserID(c *gin.Context, userID uint64) (*model.Doctor, error) {
	doctor := &model.Doctor{}
	err := repository.GetDB().WithContext(c).Model(&model.Doctor{}).Preload(clause.Associations).
		Where("user_id = ?", userID).Find(doctor).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if doctor.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return doctor, nil
}

func (ins *DoctorDal) UpdateDoctor(c *gin.Context, doctor *model.Doctor) error {
	err := repository.GetDB().WithContext(c).Table(doctor.TableName()).Updates(doctor).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) FindDoctors(c *gin.Context, currentPage, pageSize int, studyDirection, hospitalName string, professionalRank int) ([]*model.Doctor, int64, error) {
	var doctors []*model.Doctor
	// id != 0 去除默认医生 默认医生用于绑定在问诊上，当id为0时表示没有医生访问该问诊记录
	sql := repository.GetDB().WithContext(c).Model(&model.Doctor{}).Preload(clause.Associations).Where("doctors.id != 0")
	if studyDirection != "" {
		sql = sql.Where("study_direction like ?", "%"+studyDirection+"%")
	}
	if hospitalName != "" {
		sql = sql.Joins("Hospital").Clauses(clause.Like{Column: "Hospital.name", Value: "%" + hospitalName + "%"})
	}
	if professionalRank != 0 {
		sql = sql.Where("professional_rank = ?", professionalRank)
	}
	var total int64
	err := sql.Count(&total).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&doctors).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return doctors, total, nil
}

func (ins *DoctorDal) DeleteDoctor(c *gin.Context, doctorID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("doctors.id = ?", doctorID).Delete(&model.Doctor{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) ActiveDoctor(c *gin.Context, doctorID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("doctors.id = ?", doctorID).Update("is_active", true).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) FindHospitalDoctors(c *gin.Context, hospitalID uint64) ([]*model.Doctor, error) {
	var doctors []*model.Doctor
	err := repository.GetDB().WithContext(c).Model(&model.Doctor{}).Preload(clause.Associations).
		Where("hospital_id = ?", hospitalID).Find(&doctors).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return doctors, nil
}
