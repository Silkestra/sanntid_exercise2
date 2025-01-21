package main

import (
	"fmt"
	"net"

	//"sync"
)

func recieveTCP() {
	IP := "10.100.23.204"
	port := "33546"
	address := net.JoinHostPort(IP, port)

	fmt.Println("before dial: ", address)
	conn, err := net.Dial("tcp", address)
	fmt.Println("after dial")
	if err != nil {
		fmt.Println("Error")
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	for{
		fmt.Println("1")
		n, err := conn.Read(buffer)

		_, err = conn.Write([]byte("hei\x00"))

		fmt.Println("2")
		if err != nil {
			fmt.Println("Error reading")
		}
	fmt.Println("Data recieved:", string(buffer[:n-1]))
	}
}


func sendTCP(){

	IP := "10.100.23.204"
	port := "34933"
	address := net.JoinHostPort(IP, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error")
		return
	}

	defer conn.Close()

	_, err = conn.Write([]byte("hei"))
	if err != nil {
		fmt.Println("Error writing data:", err)
	}
}

func main() {
	recieveTCP()
}