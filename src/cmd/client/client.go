package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:18525")
	if err != nil {
		fmt.Println("ResolveUDPAddr fail! info:", err.Error())
		return
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("ListenUDP fail! info:", err.Error())
		return
	}
	defer udpConn.Close()

	fmt.Println("Udp dial ok!")

	len, err := udpConn.Write([]byte("Client send hello\n"))
	if err != nil{
		fmt.Println("Client send fail! info:", err.Error())
		return
	}
	fmt.Println("client write len:", len)

	buf := make([]byte, 1024)
	len, _ = udpConn.Read(buf)
	fmt.Println("client read len:", len)
	fmt.Println("client read data:", string(buf))
}
