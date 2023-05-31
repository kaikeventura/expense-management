package main

import (
	"github.com/kaikeventura/expense-management/configuration/server"
)

func main() {
	startServer()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}
