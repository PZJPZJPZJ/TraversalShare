package internal

import (
	"fmt"
	"github.com/ccding/go-stun/stun"
)

func NatTest() {
	client := stun.NewClient()
	client.SetServerAddr("stunserver.stunprotocol.org:3478")
	// 显示测试日志
	client.SetVerbose(true)

	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("NAT Type:", nat)
	fmt.Println("External IP Family:", host.Family())
	fmt.Println("External IP:", host.IP())
	fmt.Println("External Port:", host.Port())
}
