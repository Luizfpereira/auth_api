package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

type HandlerDetails struct {
	HttpMethod string
	Handler    gin.HandlerFunc
}

func NewServer(port string, router *gin.Engine) *Server {
	return &Server{
		Port:   port,
		Router: router,
	}
}

func (s *Server) Start() {
	srv := &http.Server{
		Addr:    s.Port,
		Handler: s.Router,
	}

	go func() {
		log.Println("Listening and serving on port: ", s.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	<-ctx.Done()
	log.Println("Server successfully ended!")
}
