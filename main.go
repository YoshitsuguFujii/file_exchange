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
		go server.Main(server_alive)
	} else {
		err := client.Main(os.Args)
		if err != nil {
			fmt.Printf("Error! %s", err)
		}

		return
	}

LOOP:
	for {
		select {
		case <-server_alive:
			go server.Main(server_alive)
		case <-hoge:
			break LOOP
		default:
		}
	}
}
