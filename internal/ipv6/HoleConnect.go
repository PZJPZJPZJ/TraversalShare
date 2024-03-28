package ipv6

import (
	"bufio"   // 用于读取标准输入
	"fmt"     // 用于格式化输出
	"net"     // 用于网络功能
	"os"      // 用于操作系统功能，如标准输入
	"strings" // 用于字符串处理
	"sync"    // 用于同步goroutines
)

// con6函数设置UDP连接并启动监听和发送消息的goroutines
func HoleConnect() {
	var wg sync.WaitGroup
	wg.Add(1) // 为监听goroutine添加等待

	conn, err := setupConnection() // 设置UDP连接
	if err != nil {
		fmt.Println("Error setting up connection:", err.Error())
		return
	}
	defer conn.Close() // 确保在函数结束时关闭连接

	peerAddr, err := getPeerAddress() // 获取对方的地址
	if err != nil {
		fmt.Println("Error getting peer address:", err.Error())
		return
	}

	go listenForMessages(conn, peerAddr, &wg) // 启动一个goroutine来监听消息

	sendInitialMessage(conn, peerAddr) // 发送初始消息

	wg.Wait() // 等待监听goroutine
}

// setupConnection函数创建UDP连接并返回
func setupConnection() (*net.UDPConn, error) {
	conn, err := net.ListenUDP("udp6", &net.UDPAddr{
		IP:   net.IPv6zero, // 使用IPv6地址
		Port: 0,            // 随机选择端口
		Zone: "",           // 地址范围为空
	})
	if err != nil {
		return nil, err
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("Listening on port: %d\n", localAddr.Port) // 输出监听的端口
	return conn, nil
}

// getPeerAddress函数从用户输入中获取对方的IP和端口
func getPeerAddress() (*net.UDPAddr, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter peer's IP: ") // 提示用户输入对方的IP
	peerIP, _ := reader.ReadString('\n')
	fmt.Print("Enter peer's port: ") // 提示用户输入对方的端口
	peerPort, _ := reader.ReadString('\n')
	peerIP = strings.TrimSpace(peerIP) // 去除输入字符串的空白符
	peerPort = strings.TrimSpace(peerPort)

	peerAddr, err := net.ResolveUDPAddr("udp6", "["+peerIP+"]:"+peerPort) // 解析对方的地址
	if err != nil {
		return nil, err
	}
	return peerAddr, nil
}

// sendInitialMessage函数向对方发送初始消息
func sendInitialMessage(conn *net.UDPConn, peerAddr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("hello"), peerAddr) // 发送"hello"消息
	if err != nil {
		fmt.Println("Error sending initial message:", err.Error())
	}
}

// listenForMessages函数监听UDP消息
func listenForMessages(conn *net.UDPConn, peerAddr *net.UDPAddr, wg *sync.WaitGroup) {
	defer wg.Done() // 确保在函数结束时通知WaitGroup

	buffer := make([]byte, 1024)   // 创建一个缓冲区
	connectionEstablished := false // 标记是否建立了连接

	for {
		n, addr, err := conn.ReadFromUDP(buffer) // 读取UDP消息
		if err != nil {
			fmt.Println("Error reading from UDP:", err.Error())
			continue
		}
		message := string(buffer[:n]) // 将消息转换为字符串

		// 如果收到"hello"消息且连接尚未建立，则建立连接并启动用户输入处理goroutine
		if message == "hello" && !connectionEstablished {
			connectionEstablished = true
			fmt.Println("Connection successful! You can start chatting now.")
			go handleUserInput(conn, peerAddr)
		} else if connectionEstablished {
			// 如果连接已建立，则输出对方的消息
			fmt.Printf("%s----------%s\n", message, addr.String())
		}
	}
}

// handleUserInput函数处理用户输入
func handleUserInput(conn *net.UDPConn, peerAddr *net.UDPAddr) {
	reader := bufio.NewReader(os.Stdin) // 创建一个读取器
	for {
		msg, _ := reader.ReadString('\n')                                   // 读取用户输入的消息
		_, err := conn.WriteToUDP([]byte(strings.TrimSpace(msg)), peerAddr) // 发送消息
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
	}
}
