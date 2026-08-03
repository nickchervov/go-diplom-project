package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Server *http.Server
}

func New(h http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    os.Getenv("TODO_PORT"),
			Handler: h,
		},
	}
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
