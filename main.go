package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建UDP连接到STUN服务器
	serverAddr, err := net.ResolveUDPAddr("udp", "stun.qq.com:3478")
	if err != nil {
		fmt.Println("Error resolving STUN server address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to STUN server:", err)
		return
	}
	defer conn.Close()

	// 构建STUN请求
	request := []byte{0, 1, 0, 0} // STUN Binding Request message
	conn.Write(request)

	// 接收STUN响应
	response := make([]byte, 512)
	_, err = conn.Read(response)
	if err != nil {
		fmt.Println("Error reading STUN response:", err)
		return
	}

	// 解析STUN响应以获取NAT类型
	if response[0] == 0x01 && response[1] == 0x01 {
		fmt.Println("NAT类型：Full Cone (Open Internet)")
	} else if response[0] == 0x02 && response[1] == 0x02 {
		fmt.Println("NAT类型：Restricted Cone")
	} else if response[0] == 0x03 && response[1] == 0x03 {
		fmt.Println("NAT类型：Port Restricted Cone")
	} else if response[0] == 0x04 && response[1] == 0x04 {
		fmt.Println("NAT类型：Symmetric")
	} else {
		fmt.Println("NAT类型：Unknown")
	}
}
