package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func con6() {
	var wg sync.WaitGroup
	wg.Add(1) // 为监听goroutine添加等待

	conn, err := setupConnection()
	if err != nil {
		fmt.Println("Error setting up connection:", err.Error())
		return
	}
	defer conn.Close()

	peerAddr, err := getPeerAddress()
	if err != nil {
		fmt.Println("Error getting peer address:", err.Error())
		return
	}

	go listenForMessages(conn, &wg)

	sendInitialMessage(conn, peerAddr)

	handleUserInput(conn, peerAddr)

	wg.Wait() // 等待监听goroutine
}

func setupConnection() (*net.UDPConn, error) {
	conn, err := net.ListenUDP("udp6", &net.UDPAddr{
		IP:   net.IPv6zero,
		Port: 0,
		Zone: "",
	})
	if err != nil {
		return nil, err
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("Listening on port: %d\n", localAddr.Port)
	return conn, nil
}

func getPeerAddress() (*net.UDPAddr, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter peer's IP: ")
	peerIP, _ := reader.ReadString('\n')
	fmt.Print("Enter peer's port: ")
	peerPort, _ := reader.ReadString('\n')
	peerIP = strings.TrimSpace(peerIP)
	peerPort = strings.TrimSpace(peerPort)

	peerAddr, err := net.ResolveUDPAddr("udp6", "["+peerIP+"]:"+peerPort)
	if err != nil {
		return nil, err
	}
	return peerAddr, nil
}

func sendInitialMessage(conn *net.UDPConn, peerAddr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("hello"), peerAddr)
	if err != nil {
		fmt.Println("Error sending initial message:", err.Error())
	}
}

func listenForMessages(conn *net.UDPConn, wg *sync.WaitGroup) {
	defer wg.Done()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err.Error())
			continue
		}
		message := string(buffer[:n])
		fmt.Printf("\nReceived '%s' from %s\n", message, addr.String())
		fmt.Print("Enter message to send: ")
	}
}

func handleUserInput(conn *net.UDPConn, peerAddr *net.UDPAddr) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send: ")
		msg, _ := reader.ReadString('\n')
		_, err := conn.WriteToUDP([]byte(strings.TrimSpace(msg)), peerAddr)
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
		}
	}
}
