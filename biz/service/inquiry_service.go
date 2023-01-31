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

type InquiryService struct{}

var (
	inquiryService     *InquiryService
	inquiryServiceOnce sync.Once
)

func GetInquiryService() *InquiryService {
	inquiryServiceOnce.Do(func() {
		inquiryService = &InquiryService{}
	})
	return inquiryService
}

func (ins *InquiryService) CreateInquiry(c *gin.Context, req *bo.CreateInquiryRequest) (*bo.CreateInquiryResponse, error) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	inquiry := &model.Inquiry{
		ID:                 utils.GenerateInquiryID(),
		UserID:             jwtStruct.UserID,
		ReplyDoctorID:      0, // 还没有回复 默认为0
		DiseaseName:        req.DiseaseName,
		Description:        req.Description,
		WeightHeight:       req.WeightHeight,
		HistoryOfAllergy:   req.HistoryOfAllergy,
		PastMedicalHistory: req.PastMedicalHistory,
		OtherInfo:          req.OtherInfo,
		ReplySuggestion:    "", // 还没有回复 默认为空串
	}
	err = dal.GetInquiryDal().CreateInquiry(c, inquiry)
	if err != nil {
		return nil, err
	}
	return &bo.CreateInquiryResponse{}, nil
}

func (ins *InquiryService) DeleteInquiry(c *gin.Context, req *bo.DeleteInquiryRequest) (*bo.DeleteInquiryResponse, error) {
	inquiryID, err := utils.StringToUint64(req.InquiryID)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetInquiryDal().TakeInquiryByID(c, inquiryID)
	if err != nil {
		return nil, err
	}
	err = dal.GetInquiryDal().DeleteInquiry(c, inquiryID)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteInquiryResponse{}, nil
}

func (ins *InquiryService) UploadReplySuggestion(c *gin.Context, req *bo.UploadReplySuggestionRequest) (*bo.UploadReplySuggestionResponse, error) {
	inquiryID, err := utils.StringToUint64(req.InquiryID)
	if err != nil {
		return nil, err
	}
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	doctor, err := dal.GetDoctorDal().TakeDoctorByUserID(c, jwtStruct.UserID)
	if err != nil {
		return nil, err
	}
	inquiry, err := dal.GetInquiryDal().TakeInquiryByID(c, inquiryID)
	if err != nil {
		return nil, err
	}
	inquiry.ReplyDoctorID = doctor.ID
	inquiry.ReplySuggestion = req.ReplySuggestion
	err = dal.GetInquiryDal().UpdateInquiry(c, inquiry)
	if err != nil {
		return nil, err
	}
	return &bo.UploadReplySuggestionResponse{}, nil
}

func (ins *InquiryService) FindInquiryTitles(c *gin.Context, req *bo.FindInquiryTitlesRequest) (*bo.FindInquiryTitlesResponse, error) {
	if req.ReplyStatus != common.REPLIED.Int() && req.ReplyStatus != common.NOT_REPLIED.Int() && req.ReplyStatus != common.ALL_STATUS.Int() ||
		req.CurrentPage < 0 || req.PageSize < 0 {
		return nil, common.USERINPUTERROR
	}
	inquiries, err := dal.GetInquiryDal().FindInquiries(c, req.CurrentPage, req.PageSize, common.ParseReplyStatus(req.ReplyStatus), req.DiseaseName)
	if err != nil {
		return nil, err
	}
	return &bo.FindInquiryTitlesResponse{
		Total:         len(inquiries),
		InquiryTitles: convertToInquiryTitleVOs(inquiries),
	}, nil
}

func (ins *InquiryService) FindInquiry(c *gin.Context, req *bo.FindInquiryRequest) (*bo.FindInquiryResponse, error) {
	inquiryID, err := utils.StringToUint64(req.InquiryID)
	if err != nil {
		return nil, err
	}
	inquiry, err := dal.GetInquiryDal().TakeInquiryByID(c, inquiryID)
	if err != nil {
		return nil, err
	}
	return &bo.FindInquiryResponse{Inquiry: &vo.InquiryVO{
		InquiryID:               req.InquiryID,
		UserName:                inquiry.User.Name,
		DiseaseName:             inquiry.DiseaseName,
		Description:             inquiry.Description,
		WeightHeight:            inquiry.WeightHeight,
		HistoryOfAllergy:        inquiry.HistoryOfAllergy,
		PastMedicalHistory:      inquiry.PastMedicalHistory,
		OtherInfo:               inquiry.OtherInfo,
		ReplyDoctorName:         inquiry.Doctor.User.Name,
		ReplyDoctorHospitalName: inquiry.Doctor.Hospital.Name,
		ReplySuggestion:         inquiry.ReplySuggestion,
		CreatedAt:               utils.TimeToDateString(inquiry.CreatedAt),
	}}, nil
}

func convertToInquiryTitleVOs(inquiries []*model.Inquiry) []*vo.InquiryTitleVO {
	var vos []*vo.InquiryTitleVO
	for _, inquiry := range inquiries {
		vos = append(vos, &vo.InquiryTitleVO{
			InquiryID:               utils.Uint64ToString(inquiry.ID),
			DiseaseName:             inquiry.DiseaseName,
			ReplyDoctorHospitalName: inquiry.Doctor.Hospital.Name,
			ReplyDoctorName:         inquiry.Doctor.User.Name,
		})
	}
	return vos
}
