package pkg

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func NewServer(host string, handler RequestHandler) error {
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
		fmt.Printf("client connected: %v \n", conn.RemoteAddr())
		go handleConnection(conn, handler)
	}
}

type RequestHandler interface {
	HandleRequest(request string) string
}

func handleConnection(conn net.Conn, handler RequestHandler) {
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
		output := handler.HandleRequest(string(buffer[:n]))
		writer.WriteString(output)

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Error flushing writer: %v\n", err)
			return
		}
	}
}
