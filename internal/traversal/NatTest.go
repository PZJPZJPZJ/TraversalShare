package traversal

import (
	"fmt"
	"github.com/ccding/go-stun/stun"
	"net"
	"strconv"
)

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
