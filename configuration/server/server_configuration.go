package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/expense-management/configuration/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   os.Getenv("SERVER_PORT"),
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	router := routes.ConfigurationRouter(server.server)

	log.Print("Server is running at port: ", server.port)
	log.Fatal(router.Run(":" + server.port))
}
