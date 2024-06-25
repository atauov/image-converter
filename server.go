package image_converter

import (
	"context"
	"github.com/atauov/image-converter/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *config.HTTPServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		Handler:        handler,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
