package main

import (
	"fmt"
	"gortc.io/stun"
	"log"
	"net"
)

func main() {
	// 创建一个UDP连接
	conn, err := net.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 创建一个 STUN 消息
	message := new(stun.Message)
	message.Type = stun.ClassRequest | stun.MethodBinding

	// 发送 STUN 请求
	_, err = message.WriteTo(conn)
	if err != nil {
		log.Fatal(err)
	}

	// 接收 STUN 响应
	buf := make([]byte, 1500)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// 解析 STUN 响应
	var externalIP net.IP
	var externalPort int

	var xorAddr stun.XORMappedAddress
	if err := xorAddr.GetFrom(message, stun.AttrXORMappedAddress); err != nil {
		log.Fatal(err)
	}

	externalIP = xorAddr.IP
	externalPort = xorAddr.Port

	fmt.Printf("External IP: %s\n", externalIP)
	fmt.Printf("External Port: %d\n", externalPort)
}
