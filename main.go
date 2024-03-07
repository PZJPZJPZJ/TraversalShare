//go:build windows
// +build windows

package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ccding/go-stun/stun" // 引入 go-stun 包，用于与 stun 服务器交互
)

const (
	stunServer = "stun.radiojar.com:3478" // stun 服务器的地址，可以换成其他可用的地址
)

// Client 客户端结构体，包含名称，本地地址，公网地址和套接字
type Client struct {
	name       string
	localAddr  *net.UDPAddr
	publicAddr *net.UDPAddr
	conn       *net.UDPConn
}

// NewClient 创建一个客户端，绑定一个本地 udp 端口，并向 stun 服务器发送请求，获取自己的公网地址
func NewClient(name string) (*Client, error) {
	c := &Client{name: name}
	var err error
	c.conn, err = net.ListenUDP("udp4", nil) // 随机分配一个本地 udp 端口
	if err != nil {
		return nil, err
	}
	c.localAddr = c.conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("%s: 本地地址: %s\n", c.name, c.localAddr)
	// 创建一个 stun 客户端，指定 stun 服务器的地址
	stunClient := stun.NewClient()
	stunClient.SetServerAddr(stunServer)
	// 向 stun 服务器发送请求，获取自己的 nat 类型和公网地址
	nat, host, err := stunClient.Discover()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s: NAT 类型: %s\n", c.name, nat)
	fmt.Printf("%s: 公网地址: %s\n", c.name, host)
	// 解析公网地址，得到公网 ip 和端口
	ip := net.IPv4(host.IP()[0], host.IP()[1], host.IP()[2], host.IP()[3])
	port := int(host.Port())
	c.publicAddr = &net.UDPAddr{
		IP:   ip,
		Port: port,
	}
	return c, nil
}

// 从套接字中读取数据，并打印出来
func (c *Client) read() {
	for {
		data := make([]byte, 1024)
		n, addr, err := c.conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("%s: 读取数据错误: %v\n", c.name, err)
			continue
		}
		fmt.Printf("%s: 从 %s 收到数据: %s\n", c.name, addr, data[:n])
	}
}

// 向目标地址发送数据，每隔一秒发送一次，直到程序退出
func (c *Client) write(target *net.UDPAddr) {
	for {
		// 向目标地址发送一条问候语，包含自己的名称
		data := []byte(fmt.Sprintf("Hello, I am %s", c.name))
		_, err := c.conn.WriteToUDP(data, target)
		if err != nil {
			fmt.Printf("%s: 发送数据错误: %v\n", c.name, err)
		} else {
			fmt.Printf("%s: 向 %s 发送数据: %s\n", c.name, target, data)
		}
		time.Sleep(time.Second) // 暂停一秒
	}
}

func main() {
	// 创建两个客户端，分别命名为 A 和 B
	clientA, err := NewClient("A")
	if err != nil {
		fmt.Println("创建客户端 A 错误:", err)
		return
	}
	clientB, err := NewClient("B")
	if err != nil {
		fmt.Println("创建客户端 B 错误:", err)
		return
	}
	// 让两个客户端分别输入对方的公网 ip 和端口，作为目标地址
	var ip string
	var port int
	fmt.Println("请输入客户端 A 的目标地址 (ip:port):")
	fmt.Scanln(&ip, &port)
	targetA := &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	fmt.Println("请输入客户端 B 的目标地址 (ip:port):")
	fmt.Scanln(&ip, &port)
	targetB := &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	// 让两个客户端同时开始读写数据
	go clientA.read()
	go clientB.read()
	go clientA.write(targetA)
	go clientB.write(targetB)
	// 主程序阻塞等待，直到按下回车键退出
	fmt.Println("按下回车键退出...")
	fmt.Scanln()
}
