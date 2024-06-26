package traversal

import (
	"bufio"
	"fmt"
	"github.com/ccding/go-stun/stun"
	"net"
	"os"
	"strconv"
	"time"
)

func HolePunch() {
	// 本地端口，用于与STUN服务器和对方通信
	localPort := 0 // 设置为0时，系统会自动分配端口

	// 创建本地UDP地址
	localAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("0.0.0.0:%d", localPort))
	if err != nil {
		fmt.Println("Error resolving local address:", err)
		return
	}
	// 创建UDP连接，用于通信
	conn, err := net.ListenUDP("udp4", localAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	// 延迟关闭连接，并处理错误
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 测试NAT类型
	NatTest(conn)

	// 读取对方的公网地址和端口
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter peer's external IP: ")
	peerIP, _ := reader.ReadString('\n')
	fmt.Print("Enter peer's external Port: ")
	var peerPort int
	_, err = fmt.Scanf("%d", &peerPort)
	if err != nil {
		return
	}

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

func NatTest(connection *net.UDPConn) (address string) {
	// 新建客户端
	//client := stun.NewClient()
	// 使用已创建的连接新建客户端
	client := stun.NewClientWithConnection(connection)
	// 设置STUN服务器
	client.SetServerAddr("stunserver.stunprotocol.org:3478")
	// 显示测试日志
	//client.SetVerbose(true)
	// 执行查询
	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 显示NAT类型和公网地址
	fmt.Println("本地UDP端口：", connection.LocalAddr().(*net.UDPAddr).Port)
	fmt.Println("NAT类型：", nat)
	fmt.Println("外部IP Family：", host.Family())
	fmt.Println("外部IP：", host.IP())
	fmt.Println("外部Port：", host.Port())
	return host.IP() + ":" + strconv.Itoa(int(host.Port()))
}
