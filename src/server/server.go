package server

import (
	r "DanielDDHM/sender/src/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func CreateNewServer() Server {
	return Server{
		port:   "3000",
		engine: gin.Default(),
	}
}

func (svr *Server) Run() {
	log.Print("[DanielDDHM-sender] Starting the server on port: ", svr.port)
	router := r.ConfigureApplicationRoutes(svr.engine)
	router.Run(":" + svr.port)
}
