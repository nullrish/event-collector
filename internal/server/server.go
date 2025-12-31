// Package server
package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string) *Server {
	mux := http.NewServeMux()

	registerRoutes(mux)

	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	fmt.Println("===== STARTING EVENT COLLECTOR SERVER =====")
	fmt.Println("[Startup] Listening on port:", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
