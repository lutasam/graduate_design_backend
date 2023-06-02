package utils

import (
	"bufio"
	"fmt"
	"github.com/lutasam/doctors/biz/common"
)

// GetGPT2Answer 获取人机医生的回复
func GetGPT2Answer(question string) (string, error) {
	// 如果未开启或question为空，直接返回默认回复
	if !IsConnActive() || question == "" {
		return common.ROBOTDEFAULTRESPONSE, nil
	}

	// 发送chat告诉服务器要使用智能问答功能
	_, err := fmt.Fprintf(GetSockConn(), "chat"+"\n")
	if err != nil {
		return "", common.CHATBOTERROR
	}

	// 发送消息给python机器人脚本
	_, err = fmt.Fprintf(GetSockConn(), question+"\n")
	if err != nil {
		return "", common.CHATBOTERROR
	}
	response, err := bufio.NewReader(GetSockConn()).ReadString('\n')
	if err != nil {
		return "", common.CHATBOTERROR
	}
	return response, nil
}
