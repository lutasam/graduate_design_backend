package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"github.com/lutasam/doctors/biz/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sort"
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

func (ins *DoctorDal) FindDoctors(c *gin.Context, currentPage, pageSize int, studyDirection, hospitalName string, professionalRank int, department string) ([]*model.Doctor, int64, error) {
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
	if department != "" {
		sql = sql.Joins("Department").Clauses(clause.Eq{Column: "Department.name", Value: department})
	}
	var total int64
	err := sql.Count(&total).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&doctors).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return doctors, total, nil
}

func (ins *DoctorDal) DeleteDoctor(c *gin.Context, doctorID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("id = ?", doctorID).Delete(&model.Doctor{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) DeleteDoctorByUserID(c *gin.Context, userID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("user_id = ?", userID).Delete(&model.Doctor{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) DeleteDoctorsByUserIDs(c *gin.Context, userIDs []uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("user_id in ?", userIDs).Delete(&model.Doctor{}).Error
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

func (ins *DoctorDal) FindHospitalDoctors(c *gin.Context, hospitalID uint64) ([]*model.Doctor, error) {
	var doctors []*model.Doctor
	err := repository.GetDB().WithContext(c).Model(&model.Doctor{}).Preload(clause.Associations).
		Where("hospital_id = ?", hospitalID).Find(&doctors).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return doctors, nil
}

func (ins *DoctorDal) UpdateDoctorRateScore(c *gin.Context, doctorID uint64, score float64) error {
	err := repository.GetDB().WithContext(c).Table(model.Doctor{}.TableName()).Where("id = ?", doctorID).
		Updates(map[string]interface{}{
			"rate_total_score": gorm.Expr("rate_total_score + ?", score),
			"rate_people_num":  gorm.Expr("rate_people_num + ?", 1),
		}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DoctorDal) IsUserRatedOnDoctor(c *gin.Context, userID, doctorID uint64) (bool, error) {
	isExist, err := repository.GetRedis().WithContext(c).SIsMember(c, common.DOCTOR_RATED+utils.Uint64ToString(doctorID), userID).Result()
	if err != nil {
		return false, common.REDISERROR
	}
	return isExist, nil
}

func (ins *DoctorDal) SetUserRatedOnDoctor(c *gin.Context, userID, doctorID uint64) error {
	_, err := repository.GetRedis().SAdd(c, common.DOCTOR_RATED+utils.Uint64ToString(doctorID), userID).Result()
	if err != nil {
		return common.REDISERROR
	}
	return nil
}

func (ins *DoctorDal) TakeDoctorRank(c *gin.Context, doctorID uint64, department string) (int, error) {
	var doctors []*model.Doctor
	sql := repository.GetDB().WithContext(c).Model(&model.Doctor{}).Preload(clause.Associations).Where("doctors.id != 0")
	if department != "" {
		sql = sql.Joins("Department").Clauses(clause.Eq{Column: "Department.name", Value: department})
	}
	err := sql.Find(&doctors).Error
	if err != nil {
		return 0, common.DATABASEERROR
	}
	sort.Slice(doctors, func(i, j int) bool {
		return doctors[i].RateTotalScore/float64(doctors[i].RatePeopleNum) > doctors[j].RateTotalScore/float64(doctors[j].RatePeopleNum)
	})
	rank := 1
	for _, doctor := range doctors {
		if doctor.ID == doctorID {
			return rank, nil
		}
		rank++
	}
	return 0, common.UNKNOWNERROR
}

func (ins *DoctorDal) SetDoctorReadCount(c *gin.Context, doctorID uint64) error {
	id := utils.Uint64ToString(doctorID)
	_, err := repository.GetRedis().Incr(c, id+common.DOCTOR_READ_COUNT_SUFFIX).Result()
	if err != nil {
		return common.REDISERROR
	}
	return nil
}

func (ins *DoctorDal) TakeDoctorReadCount(c *gin.Context, doctorID uint64) (int, error) {
	id := utils.Uint64ToString(doctorID)
	readCount, err := repository.GetRedis().Get(c, id+common.DOCTOR_READ_COUNT_SUFFIX).Result()
	if err != nil {
		return 0, common.REDISERROR
	}
	num, err := utils.StringToInt(readCount)
	if err != nil {
		return 0, err
	}
	return num, nil
}
