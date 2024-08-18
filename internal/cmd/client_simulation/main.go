package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/ozwin/interview-assignment-sip/internal/configs"
)

func main() {
	const clientCount = 2
	var wg sync.WaitGroup
	wg.Add(clientCount)

	addresses := []string{
		"0142e2fa3543cb32bf000100620002\n", "2\n", "014d20b404146b9a04000100620002\n", "5\n", "TUcns7plzjTV3UGwwWGRhNFzANqtf5\n",
	}

	for index := range clientCount {
		go func(index int) {
			defer wg.Done()
			simulateClient(index+1, addresses[index])
		}(index)
	}

	wg.Wait()
}

func simulateClient(id int, address string) {
	retries := configs.MaxResends
	conn, err := net.Dial("tcp", configs.ServerAddress)
	if err != nil {
		fmt.Printf("error while pinging a server %v", err.Error())
		log.Fatalln("Please check if the SIP Server is running")
	}
	defer conn.Close()
	for {

		_, err = conn.Write([]byte(address))
		if err != nil {
			if err == net.ErrClosed {
				fmt.Printf("connection is closed by server for client %v \n", conn.LocalAddr())
				break
			}
			log.Fatalf("error while sending messege to server %v", err.Error())
		}
		var response []byte
		buffer := make([]byte, 1024)
		for {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			n, err := conn.Read(buffer)
			if err != nil {
				//ignoring error for now, assuming it throws deadline reached error
				break
			}
			response = append(response, buffer[:n]...)
		}
		fmt.Printf("client: %v address: %v response %v\n", conn.LocalAddr(), address, string(response))
		if id%2 == 0 {
			fmt.Printf("client %v is going to sleep\n", conn.LocalAddr())
			time.Sleep(configs.SleepTimeForSimulation)
			return
		} else if retries == 0 {
			return
		} else {
			fmt.Printf("resending request for client: %v \n", conn.LocalAddr())
			retries--
		}

	}
}
