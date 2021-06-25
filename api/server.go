package api

import (
	"net/http"

	cache "github.com/yobol/go-cache/core"
)

func New(c cache.Cache) *Server {
	return &Server{c}
}

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":10615", nil)
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
