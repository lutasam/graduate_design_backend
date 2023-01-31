package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"gorm.io/gorm/clause"
	"sync"
)

// 要改
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
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Preload(clause.Associations).
		Where("id = ?", doctorID).Find(doctor).Error
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
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Preload(clause.Associations).
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

func (ins *DoctorDal) FindDoctors(c *gin.Context, currentPage, pageSize int, name, hospitalName, departmentName string) ([]*model.Doctor, error) {
	var doctors []*model.Doctor
	// id != 0 去除默认医生
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("id != 0").
		Joins("User").Clauses(clause.Like{Column: "User.name", Value: "%" + name + "%"}).
		Joins("Department").Clauses(clause.Like{Column: "Hospital.name", Value: "%" + hospitalName + "%"}).
		Joins("Hospital").Clauses(clause.Like{Column: "Department.name", Value: "%" + departmentName + "%"}).
		Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&doctors).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return doctors, nil
}

func (ins *DoctorDal) DeleteDoctor(c *gin.Context, doctorID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("id = ?", doctorID).Delete(&model.Doctor{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) ActiveDoctor(c *gin.Context, doctorID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("id = ?", doctorID).Update("is_active", true).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
