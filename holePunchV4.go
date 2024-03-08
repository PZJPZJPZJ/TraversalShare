package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ccding/go-stun/stun"
)

func con4() {
	// 本地端口，用于与STUN服务器和对方通信
	localPort := 0 // 设置为0时，系统会自动分配端口

	// 创建本地UDP地址
	localAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("0.0.0.0:%d", localPort))
	if err != nil {
		fmt.Println("Error resolving local address:", err)
		return
	}

	// 创建UDP连接，用于与STUN服务器通信
	conn, err := net.ListenUDP("udp4", localAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	// 打印本地端口
	localPort = conn.LocalAddr().(*net.UDPAddr).Port
	fmt.Printf("Local UDP port before contacting STUN servers: %d\n", localPort)

	// 使用STUN库查询第一个公网地址和端口
	client := stun.NewClientWithConnection(conn)
	client.SetServerAddr("stun.radiojar.com:3478")

	nat1, host1, err := client.Discover()
	if err != nil {
		fmt.Println("Error discovering NAT type from first STUN server:", err)
		return
	}

	// 打印本地端口
	fmt.Printf("Local UDP port after contacting first STUN server: %d\n", localPort)

	// 使用STUN库查询第二个公网地址和端口
	client.SetServerAddr("stun.miwifi.com:3478")
	nat2, host2, err := client.Discover()
	if err != nil {
		fmt.Println("Error discovering NAT type from second STUN server:", err)
		return
	}

	// 检查两个外部端口是否一致
	if host1.Port() != host2.Port() {
		fmt.Println("The external ports are not consistent.")
		return
	}

	// 显示NAT类型和公网地址
	fmt.Printf("NAT Type from first STUN server: %s\n", nat1)
	fmt.Printf("External IP from first STUN server: %s\n", host1.IP())
	fmt.Printf("External Port from first STUN server: %d\n", host1.Port())

	fmt.Printf("NAT Type from second STUN server: %s\n", nat2)
	fmt.Printf("External IP from second STUN server: %s\n", host2.IP())
	fmt.Printf("External Port from second STUN server: %d\n", host2.Port())

	// 读取对方的公网地址和端口
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter peer's external IP: ")
	peerIP, _ := reader.ReadString('\n')
	fmt.Print("Enter peer's external Port: ")
	var peerPort int
	fmt.Scanf("%d", &peerPort)

	// 设置对方的地址
	peerAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", peerIP, peerPort))
	if err != nil {
		fmt.Println("Error resolving peer address:", err)
		return
	}

	// 使用相同的本地端口与对方通信
	_, err = conn.WriteTo([]byte("Hello from NAT!"), peerAddr)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// 接收对方的消息
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		return
	}

	// 打印接收到的消息
	fmt.Printf("Received message: %s\n", string(buffer[:n]))

	// 发送心跳包保持连接
	go func() {
		for {
			time.Sleep(30 * time.Second)
			_, err = conn.WriteTo([]byte("Heartbeat"), peerAddr)
			if err != nil {
				fmt.Println("Error sending heartbeat:", err)
				return
			}
		}
	}()

	// 持续监听该端口收到的消息，并打印在控制台上
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(buffer[:n]))
	}
}
