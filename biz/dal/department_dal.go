package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
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
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("id = ?", departmentID).Find(department).Error
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
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("name = ?", departmentName).Find(department).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if department.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return department, nil
}
