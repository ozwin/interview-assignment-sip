package main

import (
	"github.com/ozwin/interview-assignment-sip/internal/app/services"
	"github.com/ozwin/interview-assignment-sip/internal/configs"
	"github.com/ozwin/interview-assignment-sip/internal/pkg"
)

func main() {
	service := services.NewTransactionService()
	pkg.NewServer(configs.ServerAddress, service)
}
