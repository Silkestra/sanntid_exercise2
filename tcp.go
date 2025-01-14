package main

import (
	"fmt"
	"net"
)


func recieveTCP(){
	address := "10.100.23.204:34933"

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error")
		return 
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listener")
		return 
	}
	defer listener.Close()


	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error reading")
			continue
		}
	go handleConnection(conn)
	}
}


func handleConnection(conn *net.TCPConn){
	defer conn.Close()

	fmt.Printf("Recieved connection from %s\n", conn.RemoteAddr())
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading")
		return 
	}

	fmt.Printf("%d bytes : %s \n", n, string(buffer[:n]))
}


func main(){
	recieveTCP() 
}
