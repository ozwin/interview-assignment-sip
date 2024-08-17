package pkg

import (
	"net"
	"testing"
	"time"
)

type DummyHandler struct{}

func (dh *DummyHandler) HandleRequest(request string) string {
	return request + " back"
}

func Test_handleConnection(t *testing.T) {
	server, client := net.Pipe()
	defer client.Close()
	defer server.Close()

	go handleConnection(server, &DummyHandler{})

	client.Write([]byte("Hello"))
	// Read the response from the server.
	var responseData []byte
	buffer := make([]byte, 1024)
	for {
		//keeping the timer for now , not sure about the scenario about EOF token
		client.SetReadDeadline(time.Now().Add(3 * time.Second))
		n, err := client.Read(buffer)
		if err != nil {
			break
		}
		responseData = append(responseData, buffer[:n]...)
	}

	expectedResponse := "Hello back"
	response := string(responseData)
	if response != expectedResponse {
		t.Errorf("expected response %q, got %q", expectedResponse, response)
	}
}
