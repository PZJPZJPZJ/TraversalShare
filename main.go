//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/pion/stun"
)

// stunRequest 发送 STUN 请求到指定的 STUN 服务器，并输出解析后的 XOR-MAPPED-ADDRESS。
func stunRequest(serverURL string) {
	// 解析 STUN URI
	u, err := stun.ParseURI("stun:" + serverURL)
	if err != nil {
		panic(err)
	}

	// 创建到 STUN 服务器的连接
	c, err := stun.DialURI(u, &stun.DialConfig{})
	if err != nil {
		panic(err)
	}

	// 构建带有随机事务 ID 的绑定请求
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	// 发送请求到 STUN 服务器，等待响应消息
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}

		// 从消息中解码 XOR-MAPPED-ADDRESS 属性
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		fmt.Println(xorAddr.IP, ":", xorAddr.Port)
	}); err != nil {
		panic(err)
	}
}

func main() {
	// 向两个不同的 STUN 服务器发送请求并输出结果
	stunRequest("stun.radiojar.com:3478")
	stunRequest("stun.miwifi.com:3478")
}
