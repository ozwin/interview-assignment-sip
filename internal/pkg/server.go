package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/ozwin/interview-assignment-sip/internal/configs"
)

func NewServer(host string, handler RequestHandler) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(fmt.Sprintf("failed to stat the server: %v", err))
	}
	defer listener.Close()
	fmt.Printf("server listening at: %v\n", host)
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
	reader := bufio.NewReader(conn)
	// buffer := make([]byte, 1024)
	for {
		conn.SetReadDeadline(time.Now().Add(configs.ConnectionTimeout))
		//keep this for now
		// n, err := conn.Read(buffer)
		payload, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Printf("client is inactive for more than 10 seconds, closing connection for %v \n", conn.RemoteAddr())
			} else if errors.Is(err, io.EOF) {
				fmt.Printf("connection closed by the client %v\n", conn.RemoteAddr())
			} else {
				fmt.Printf("unknown error %v\n", err)
			}
			return
		}
		response := handler.HandleRequest(payload)

		if _, err = writer.WriteString(response); err != nil {
			fmt.Printf("error while writing back to the client: %v\n", err)
			return
		}

		if err = writer.Flush(); err != nil {
			fmt.Printf("error flushing writer: %v\n", err)
			return
		}
	}
}
