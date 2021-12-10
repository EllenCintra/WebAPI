package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine //preciso para iniciar o server
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
