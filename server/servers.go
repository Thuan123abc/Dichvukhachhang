package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func NewHTTPServer(port string) *Server {
	deployTime := time.Now()
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins: true,
	}))
	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"deployTime": deployTime,
			"timestamp":  time.Now(),
		})
	})
	server := &Server{
		port:   port,
		engine: engine,
	}
	server.RegisterRoutes()
	return server
}

func (s *Server) Serve() error {
	return s.engine.Run(fmt.Sprintf(":%s", s.port))
}
