package main

import (
	"fmt"

	"github.com/ozwin/interview-assignment-sip/internal/pkg"
)

const serverAddres = "localhost:8080"

func main() {
	pkg.NewServer(serverAddres, Test)
}

func Test(messege string) string {
	fmt.Println(messege)
	return messege + " back "
}
