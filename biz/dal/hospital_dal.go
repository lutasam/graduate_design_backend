package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"sync"
)

type HospitalDal struct{}

var (
	hospitalDal     *HospitalDal
	hospitalDalOnce sync.Once
)

func GetHospitalDal() *HospitalDal {
	hospitalDalOnce.Do(func() {
		hospitalDal = &HospitalDal{}
	})
	return hospitalDal
}

func (ins *HospitalDal) TakeHospitalByID(c *gin.Context, hospitalID uint64) (*model.Hospital, error) {
	hospital := &model.Hospital{}
	err := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("id = ?", hospitalID).Find(hospital).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if hospital.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return hospital, nil
}

func (ins *HospitalDal) TakeHospitalByName(c *gin.Context, hospitalName string) (*model.Hospital, error) {
	hospital := &model.Hospital{}
	err := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("name = ?", hospitalName).Find(hospital).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if hospital.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return hospital, nil
}

func (ins *HospitalDal) FindHospitals(c *gin.Context, currentPage, pageSize int, hospitalName, city string, hospitalRank int) ([]*model.Hospital, int64, error) {
	var hospitals []*model.Hospital
	sql := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("id != ?", 0)
	if hospitalName != "" {
		sql = sql.Where("name like ?", "%"+hospitalName+"%")
	}
	if city != "" {
		sql = sql.Where("city = ?", city)
	}
	if hospitalRank != 0 {
		sql = sql.Where("hospital_rank = ?", hospitalRank)
	}
	var total int64
	err := sql.Count(&total).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&hospitals).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return hospitals, total, nil
}
