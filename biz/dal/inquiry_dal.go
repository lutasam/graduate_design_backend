package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"github.com/lutasam/doctors/biz/repository"
	"github.com/lutasam/doctors/biz/utils"
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
	err := repository.GetDB().WithContext(c).Model(&model.Inquiry{}).Preload(clause.Associations).
		Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital").
		Where("inquiries.id = ?", inquiryID).Find(inquiry).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if inquiry.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return inquiry, nil
}

func (ins *InquiryDal) DeleteInquiry(c *gin.Context, inquiryID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Inquiry{}.TableName()).Where("inquiries.id = ?", inquiryID).
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

func (ins *InquiryDal) FindInquiries(c *gin.Context, currentPage, pageSize int, replyStatus common.ReplyStatus, diseaseName string) ([]*model.Inquiry, int64, error) {
	// new version use cosine similarity
	// 如果diseaseName不为空, 则可以进行语义查询
	if diseaseName != "" {
		// 1. 获取inquiryES, 并获取对应的 inquiryIDs 数组
		inquiryES, err := utils.SearchInquiriesInES(diseaseName)
		if err != nil {
			return nil, 0, err
		}
		var inquiryIDs []uint64
		for _, es := range inquiryES {
			inquiryIDs = append(inquiryIDs, es.InquiryID)
		}
		// 2. 从db中获取对应id的inquiry, 并返回
		var inquiries []*model.Inquiry
		// 通过order by field id 按照es返回的id顺序返回结果
		sql := repository.GetDB().WithContext(c).Model(&model.Inquiry{}).Preload(clause.Associations).
			Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital").
			Where("id in ?", inquiryIDs).Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{inquiryIDs}, WithoutParentheses: true},
		})
		if replyStatus != common.ALL_STATUS {
			if replyStatus == common.REPLIED {
				sql = sql.Where("reply_doctor_id != ?", 0) // 存在回复则该条记录回复医生的id不为0
			} else {
				sql = sql.Where("reply_doctor_id = ?", 0) // 存在回复则该条记录回复医生的id为0
			}
		}
		err = sql.Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&inquiries).Error
		if err != nil {
			return nil, 0, common.DATABASEERROR
		}
		return inquiries, int64(len(inquiryES)), nil
	} else { // 否则, 走普通的db查询
		var inquiries []*model.Inquiry
		sql := repository.GetDB().WithContext(c).Model(&model.Inquiry{}).Preload(clause.Associations).
			Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital")
		if replyStatus != common.ALL_STATUS {
			if replyStatus == common.REPLIED {
				sql = sql.Where("reply_doctor_id != ?", 0) // 存在回复则该条记录回复医生的id不为0
			} else {
				sql = sql.Where("reply_doctor_id = ?", 0) // 存在回复则该条记录回复医生的id为0
			}
		}
		var total int64
		err := sql.Count(&total).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&inquiries).Error
		if err != nil {
			return nil, 0, common.DATABASEERROR
		}
		return inquiries, total, nil
	}

	// old version just search keyword in db
	//var inquiries []*model.Inquiry
	//sql := repository.GetDB().WithContext(c).Model(&model.Inquiry{}).Preload(clause.Associations).
	//	Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital")
	//if replyStatus != common.ALL_STATUS {
	//	if replyStatus == common.REPLIED {
	//		sql = sql.Where("reply_doctor_id != ?", 0) // 存在回复则该条记录回复医生的id不为0
	//	} else {
	//		sql = sql.Where("reply_doctor_id = ?", 0) // 存在回复则该条记录回复医生的id为0
	//	}
	//}
	//if diseaseName != "" {
	//	sql = sql.Where("disease_name like ?", "%"+diseaseName+"%")
	//}
	//var total int64
	//err := sql.Count(&total).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&inquiries).Error
	//if err != nil {
	//	return nil, 0, common.DATABASEERROR
	//}
	//return inquiries, total, nil

}

func (ins *InquiryDal) FindDoctorInquiries(c *gin.Context, doctorID uint64) ([]*model.Inquiry, error) {
	var inquiries []*model.Inquiry
	err := repository.GetDB().WithContext(c).Model(&model.Inquiry{}).Preload(clause.Associations).
		Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital").
		Where("reply_doctor_id = ?", doctorID).Find(&inquiries).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return inquiries, nil
}

func (ins *InquiryDal) FindUserInquiries(c *gin.Context, userID uint64) ([]*model.Inquiry, error) {
	var inquiries []*model.Inquiry
	err := repository.GetDB().WithContext(c).Model(&model.Inquiry{}).Preload(clause.Associations).
		Preload("Doctor.User").Preload("Doctor.Department").Preload("Doctor.Hospital").
		Where("user_id = ?", userID).Find(&inquiries).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return inquiries, nil
}
