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
	Port     string
	Handlers map[string]HandlerDetails
	Router   *gin.Engine
}

type HandlerDetails struct {
	HttpMethod string
	Handler    gin.HandlerFunc
}

func NewServer(port string) *Server {
	return &Server{
		Port:     port,
		Handlers: make(map[string]HandlerDetails),
		Router:   gin.Default(),
	}
}

func (s *Server) AddHandler(path, httpMethod string, handler gin.HandlerFunc) {
	s.Handlers[path] = HandlerDetails{HttpMethod: httpMethod, Handler: handler}
}

func (s *Server) Start() {
	for path, handlerDetails := range s.Handlers {
		s.Router.Handle(handlerDetails.HttpMethod, path, handlerDetails.Handler)
	}

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
