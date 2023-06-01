package main

import (
	"github.com/kaikeventura/expense-management/configuration/container"
	"github.com/kaikeventura/expense-management/configuration/database"
	"github.com/kaikeventura/expense-management/configuration/server"
)

func main() {
	startDatabase()
	startDependencyInjection()
	startServer()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}

func startDatabase() {
	database.RunDatabase()
}

func startDependencyInjection() {
	container.BuildDependencyInjection()
}
