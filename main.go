package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/poximy/crypto-price-tracker/pages"
)

func main() {
	s := NewServer()
	s.MountMiddleware()
	s.MountHandlers()

	err := http.ListenAndServe(port(), s.Router)
	if err != nil {
		panic(err)
	}
}

type Server struct {
	Router *chi.Mux
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountHandlers() {
	s.Router.Mount("/", pages.NewRoute())
}

func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func port() string {
	portNum := os.Getenv("PORT")
	if portNum == "" {
		portNum = "8080" // Default port if not specified
	}
	return ":" + portNum
}
