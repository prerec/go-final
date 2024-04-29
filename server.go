package todo

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	log.Printf("Starting server on http://localhost:%s/", port)
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.httpServer.Handler.ServeHTTP(w, r)
}
