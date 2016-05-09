package server

import (
	"fmt"
	"net"
)

const UdpPort = ":8889"

func RunUdpServ() {
	fmt.Println("### Server(Udp) Start on" + UdpPort + " ###")

	addr, err := net.ResolveUDPAddr("udp", UdpPort)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		rlen, remote, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Error: %v\n", err)
		}

		s := string(buf[:rlen])

		fmt.Println("Receive [%v]: %v\n", remote, s)

		s = "Hello! " + s

		rlen, err = conn.WriteToUDP([]byte(s), remote)

		if err != nil {
			fmt.Println("Receive Error [%v]: %v\n", remote, s)
		}

		fmt.Println("Send [%v]: %v\n", remote, s)
	}
}
