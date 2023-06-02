package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/model"
	"strings"
)

// SearchInquiriesInES find inquiries in elasticsearch by searchText's cosine similarity
func SearchInquiriesInES(searchText string) ([]*model.InquiryES, error) {
	if !IsConnActive() {
		return nil, common.ELASTICSEARCHERROR
	}

	// 发送search告诉服务器要使用语义搜索功能
	_, err := fmt.Fprintf(GetSockConn(), "search"+"\n")
	if err != nil {
		return nil, common.ELASTICSEARCHERROR
	}

	// 1. 通过python文件计算余弦相似度并返回es搜索结果
	_, err = fmt.Fprintf(GetSockConn(), searchText)
	if err != nil {
		return nil, common.ELASTICSEARCHERROR
	}
	response, err := bufio.NewReader(GetSockConn()).ReadString('\n')
	if err != nil {
		return nil, common.ELASTICSEARCHERROR
	}

	// 2. python输出结果转换为正确的json格式
	jsonStr := strings.Replace(response, "'", "\"", -1)

	// 3. json反序列化到对象
	var inquiries []*model.InquiryES
	err = json.Unmarshal([]byte(jsonStr), &inquiries)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return inquiries, nil
}

// InsertInquiryToES add inquiry into elasticsearch
func InsertInquiryToES(inquiry *model.InquiryES) error {
	if !IsConnActive() {
		return common.ELASTICSEARCHERROR
	}
	_, err := fmt.Fprintf(GetSockConn(), "add"+"\n")
	if err != nil {
		return common.ELASTICSEARCHERROR
	}
	bytes, err := json.Marshal(inquiry)
	if err != nil {
		return common.UNKNOWNERROR
	}
	_, err = fmt.Fprintf(GetSockConn(), string(bytes))
	if err != nil {
		return common.ELASTICSEARCHERROR
	}
	response, err := bufio.NewReader(GetSockConn()).ReadString('\n')
	if err != nil || response != "ok\n" {
		return common.ELASTICSEARCHERROR
	}
	return nil
}
