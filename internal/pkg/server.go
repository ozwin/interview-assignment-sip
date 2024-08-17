package pkg

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func NewServer(host string, requestHandler RequestHandler) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(fmt.Sprintf("failed to stat the server: %v", err))
	}
	defer listener.Close()
	fmt.Printf("server listening at: %v", host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %v", err)
			continue
		}
		fmt.Println("client connected")
		go handleConnection(conn, requestHandler)
	}
}

type RequestHandler func(request string) string

func handleConnection(conn net.Conn, responseHandler RequestHandler) {
	defer conn.Close()
	writer := bufio.NewWriter(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))

		//keep this for now
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Printf("error occured: %v", err.Error())
				continue
			} else {
				break
			}
		}
		output := responseHandler(string(buffer[:n]))
		writer.WriteString(output)

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Error flushing writer: %v\n", err)
			return
		}
	}
}
