package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"gorm.io/gorm/clause"
	"sync"
)

type InquiryDal struct{}

var (
	inquiryDal     *InquiryDal
	inquiryDalOnce sync.Once
)

func GetInquiryDal() *InquiryDal {
	inquiryDalOnce.Do(func() {
		inquiryDal = &InquiryDal{}
	})
	return inquiryDal
}

func (ins *InquiryDal) CreateInquiry(c *gin.Context, inquiry *model.Inquiry) error {
	err := repository.GetDB().WithContext(c).Table(inquiry.TableName()).Create(inquiry).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *InquiryDal) TakeInquiryByID(c *gin.Context, inquiryID uint64) (*model.Inquiry, error) {
	inquiry := &model.Inquiry{}
	err := repository.GetDB().WithContext(c).Table(inquiry.TableName()).Preload(clause.Associations).
		Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital").
		Where("id = ?", inquiryID).Find(inquiry).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if inquiry.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return inquiry, nil
}

func (ins *InquiryDal) DeleteInquiry(c *gin.Context, inquiryID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Inquiry{}.TableName()).Where("id = ?", inquiryID).
		Delete(&model.Inquiry{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *InquiryDal) UpdateInquiry(c *gin.Context, inquiry *model.Inquiry) error {
	err := repository.GetDB().WithContext(c).Table(inquiry.TableName()).Updates(inquiry).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *InquiryDal) FindInquiries(c *gin.Context, currentPage, pageSize int, replyStatus common.ReplyStatus, diseaseName string) ([]*model.Inquiry, error) {
	var inquiries []*model.Inquiry
	sql := repository.GetDB().WithContext(c).Table(model.Inquiry{}.TableName()).Preload(clause.Associations).
		Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital")
	if replyStatus != common.ALL_STATUS {
		if replyStatus == common.REPLIED {
			sql = sql.Where("reply_doctor_id != ?", 0) // 存在回复则该条记录回复医生的id不为0
		} else {
			sql = sql.Where("reply_doctor_id = ?", 0) // 存在回复则该条记录回复医生的id为0
		}
	}
	if diseaseName != "" {
		sql = sql.Where("disease_name like ?", "%"+diseaseName+"%")
	}
	err := sql.Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&inquiries).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return inquiries, nil
}
