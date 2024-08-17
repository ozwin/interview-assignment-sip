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
	const clientCount = 5
	var wg sync.WaitGroup
	wg.Add(clientCount)

	addresses := []string{
		"0142e2fa3543cb32bf000100620002", "014d20b404146b9a04000100620002", "4", "5", "TUcns7plzjTV3UGwwWGRhNFzANqtf5",
	}

	for index := range clientCount {
		go func(index int) {
			defer wg.Done()
			simulateClient(index, addresses[index])
		}(index)
	}

	wg.Wait()
}

func simulateClient(id int, address string) {
	conn, err := net.Dial("tcp", configs.ServerAddress)
	if err != nil {
		fmt.Printf("error while pinging a server %v", err.Error())
		log.Fatalln("Please check if the SIP Server is running")
	}
	defer conn.Close()
	_, err = conn.Write([]byte(address))
	if err != nil {
		fmt.Printf("error while sending messege to server %v", err.Error())
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
	fmt.Printf("client: %v address: %v response %v\n", id, address, string(response))
}
