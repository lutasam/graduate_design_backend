package utils

import (
	"fmt"
	"net"
)

var (
	conn     net.Conn
	isActive bool
)

// 初始化socket连接，连接到python的智能机器人后台
func init() {
	var err error
	conn, err = net.Dial("tcp", GetConfigString("sock.host")+":"+GetConfigString("sock.port"))
	if err != nil {
		fmt.Println("python socket未开启")
		isActive = false
	} else {
		fmt.Println("python socket已开启")
		isActive = true
	}
}

func GetSockConn() net.Conn {
	return conn
}

func IsConnActive() bool {
	return isActive
}
