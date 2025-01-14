package main

import (
	"fmt"
	"net"
	"time"
)

//39 bytes, 10.100.23.204:55555: Hello from UDP server at 10.100.23.204!

func recieveUDP() {

	addr := &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 20013,
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error")
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading")
			continue
		}
		fmt.Printf("%d bytes, %s: %s \n", n, remoteAddr, string(buffer[:n]))
	}

	fmt.Println("Hello world")

}

func sendUDP() {
	addr := &net.UDPAddr{
		IP:   net.ParseIP("10.100.23.204"),
		Port: 20013,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error")
	}
	defer conn.Close()
	fmt.Printf("Send")
	_, err = conn.Write([]byte("hei"))
	if err != nil {
		fmt.Println("Error sending")
	}
	time.Sleep(1000 * time.Millisecond)

}

func main() {

	go sendUDP()
	recieveUDP()

}
