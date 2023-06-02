package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"github.com/lutasam/doctors/biz/utils"
	"gorm.io/gorm"
	"sort"
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

func (ins *HospitalDal) FindHospitals(c *gin.Context, currentPage, pageSize int, hospitalName, city string, hospitalRank int, characteristic string) ([]*model.Hospital, int64, error) {
	var hospitals []*model.Hospital
	sql := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("id != ?", 0)
	if hospitalName != "" {
		sql = sql.Where("name like ?", "%"+hospitalName+"%")
	}
	if city != "" {
		sql = sql.Where("city like ?", "%"+city+"%") // 可能用户只输入省份没输入城市，所以用like搜索
	}
	if hospitalRank != 0 {
		sql = sql.Where("hospital_rank = ?", hospitalRank)
	}
	if characteristic != "" {
		sql = sql.Where("characteristic = ?", characteristic)
	}
	var total int64
	err := sql.Count(&total).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&hospitals).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return hospitals, total, nil
}

func (ins *HospitalDal) CreateHospital(c *gin.Context, hospital *model.Hospital) error {
	err := repository.GetDB().WithContext(c).Table(hospital.TableName()).Create(hospital).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *HospitalDal) UpdateHospitalInfo(c *gin.Context, hospital *model.Hospital) error {
	err := repository.GetDB().WithContext(c).Table(hospital.TableName()).Updates(hospital).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *HospitalDal) DeleteHospital(c *gin.Context, hospitalID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("id = ?", hospitalID).Delete(&model.Hospital{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *HospitalDal) UpdateHospitalRateScore(c *gin.Context, hospitalID uint64, score float64) error {
	err := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("id = ?", hospitalID).
		Updates(map[string]interface{}{
			"rate_total_score": gorm.Expr("rate_total_score + ?", score),
			"rate_people_num":  gorm.Expr("rate_people_num + ?", 1),
		}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *HospitalDal) IsUserRatedOnHospital(c *gin.Context, userID, hospitalID uint64) (bool, error) {
	isExist, err := repository.GetRedis().WithContext(c).SIsMember(c, common.HOSPITAL_RATED+utils.Uint64ToString(hospitalID), userID).Result()
	if err != nil {
		return false, common.REDISERROR
	}
	return isExist, nil
}

func (ins *HospitalDal) SetUserRatedOnHospital(c *gin.Context, userID, hospitalID uint64) error {
	_, err := repository.GetRedis().SAdd(c, common.HOSPITAL_RATED+utils.Uint64ToString(hospitalID), userID).Result()
	if err != nil {
		return common.REDISERROR
	}
	return nil
}

func (ins *HospitalDal) TakeHospitalRank(c *gin.Context, hospitalID uint64, area string) (int, error) {
	var hospitals []*model.Hospital
	sql := repository.GetDB().WithContext(c).Table(model.Hospital{}.TableName()).Where("id != ?", 0)
	if area != "" {
		sql = sql.Where("city = ?", area)
	}
	err := sql.Find(&hospitals).Error
	if err != nil {
		return 0, common.DATABASEERROR
	}
	sort.Slice(hospitals, func(i, j int) bool {
		return hospitals[i].RateTotalScore/float64(hospitals[i].RatePeopleNum) > hospitals[j].RateTotalScore/float64(hospitals[j].RatePeopleNum)
	})
	rank := 1
	for _, hospital := range hospitals {
		if hospital.ID == hospitalID {
			return rank, nil
		}
		rank++
	}
	return 0, common.UNKNOWNERROR
}

func (ins *HospitalDal) SetHospitalReadCount(c *gin.Context, hospitalID uint64) error {
	id := utils.Uint64ToString(hospitalID)
	_, err := repository.GetRedis().Incr(c, id+common.HOSPITAL_READ_COUNT_SUFFIX).Result()
	if err != nil {
		return common.REDISERROR
	}
	return nil
}

func (ins *HospitalDal) TakeHospitalReadCount(c *gin.Context, hospitalID uint64) (int, error) {
	id := utils.Uint64ToString(hospitalID)
	readCount, err := repository.GetRedis().Get(c, id+common.HOSPITAL_READ_COUNT_SUFFIX).Result()
	if err != nil {
		return 0, common.REDISERROR
	}
	num, err := utils.StringToInt(readCount)
	if err != nil {
		return 0, err
	}
	return num, nil
}
