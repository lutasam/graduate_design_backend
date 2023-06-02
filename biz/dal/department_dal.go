package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"gorm.io/gorm/clause"
	"sync"
)

type DepartmentDal struct{}

var (
	departmentDal     *DepartmentDal
	departmentDalOnce sync.Once
)

func GetDepartmentDal() *DepartmentDal {
	departmentDalOnce.Do(func() {
		departmentDal = &DepartmentDal{}
	})
	return departmentDal
}

func (ins *DepartmentDal) TakeDepartmentByID(c *gin.Context, departmentID uint64) (*model.Department, error) {
	department := &model.Department{}
	err := repository.GetDB().WithContext(c).Model(&model.Department{}).Preload(clause.Associations).Where("departments.id = ?", departmentID).Find(department).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if department.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return department, nil
}

func (ins *DepartmentDal) TakeDepartmentByName(c *gin.Context, departmentName string) (*model.Department, error) {
	department := &model.Department{}
	err := repository.GetDB().WithContext(c).Model(&model.Department{}).Preload(clause.Associations).Where("departments.name = ?", departmentName).Find(department).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if department.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return department, nil
}

func (ins *DepartmentDal) FindHospitalDepartments(c *gin.Context, hospitalID uint64) ([]*model.Department, error) {
	var departments []*model.Department
	err := repository.GetDB().WithContext(c).Model(&model.Department{}).Preload(clause.Associations).
		Where("hospital_id = ?", hospitalID).Find(&departments).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return departments, nil
}

func (ins *DepartmentDal) CreateDepartment(c *gin.Context, department *model.Department) error {
	err := repository.GetDB().WithContext(c).Table(department.TableName()).Create(department).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DepartmentDal) DeleteDepartment(c *gin.Context, departmentID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("id = ?", departmentID).Delete(&model.Department{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *DepartmentDal) DeleteDepartmentsByHospitalID(c *gin.Context, hospitalID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("hospital_id = ?", hospitalID).Delete(&model.Department{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
