package wireguard

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func CreateTunnel() {
	// 创建 WireGuard 控制器
	wg, err := wgctrl.New()
	if err != nil {
		log.Fatalf("无法创建 WireGuard 控制器: %v", err)
	}
	defer wg.Close()

	// 生成私钥
	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Fatalf("无法生成私钥: %v", err)
	}

	// 生成公钥
	pubKey := key.PublicKey()

	// 配置 WireGuard 设备接口
	cfg := wgtypes.Config{
		PrivateKey:   &key,
		ListenPort:   new(int), // 自动选择端口
		ReplacePeers: true,
	}

	// 创建 WireGuard 设备接口
	ifName := "wg0"
	err = wg.ConfigureDevice(ifName, cfg)
	if err != nil {
		log.Fatalf("无法配置 WireGuard 设备接口 %s: %v", ifName, err)
	}

	fmt.Printf("私钥: %s\n", key.String())
	fmt.Printf("公钥: %s\n", pubKey.String())

	// 假设对端的公钥和端点地址已知
	peerPubKeyStr := "对端公钥"
	peerPubKey, err := wgtypes.ParseKey(peerPubKeyStr)
	if err != nil {
		log.Fatalf("无法解析对端公钥: %v", err)
	}

	peerEndpointStr := "对端地址:端口"
	peerEndpoint, err := net.ResolveUDPAddr("udp", peerEndpointStr)
	if err != nil {
		log.Fatalf("无法解析对端端点地址: %v", err)
	}

	// 配置对端
	peerCfg := wgtypes.PeerConfig{
		PublicKey:                   peerPubKey,
		Endpoint:                    peerEndpoint,
		ReplaceAllowedIPs:           true,
		AllowedIPs:                  []net.IPNet{{IP: net.IPv4(192, 168, 1, 2), Mask: net.CIDRMask(32, 32)}}, // 允许的 IP 地址范围
		PersistentKeepaliveInterval: new(time.Duration),                                                      // 心跳包间隔
	}

	// 添加对端到 WireGuard 设备接口
	err = wg.ConfigureDevice(ifName, wgtypes.Config{Peers: []wgtypes.PeerConfig{peerCfg}})
	if err != nil {
		log.Fatalf("无法添加对端: %v", err)
	}

	fmt.Println("P2P 隧道创建成功")
}
