package main

import (
	"./client"
	"./server"
	"fmt"
	"os"
	"time"
)

const IntervalSec = 5 * time.Second

var p = fmt.Println

func main() {
	hoge := make(chan bool)
	server_alive := make(chan bool)

	if len(os.Args) == 1 || os.Args[1] == "--server" {
		go server.Run(server_alive)
		go server.RunUdpServ()
	} else {
		err := client.Main(os.Args)
		if err != nil {
			fmt.Printf("Error! %s", err)
		}

		return
	}

	// 現状無意味なループ
LOOP:
	for {
		select {
		case <-server_alive:
			go server.Run(server_alive)
		case <-hoge:
			break LOOP
		default:
		}
	}
}
