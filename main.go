package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/poximy/crypto-price-tracker/utils"
	"golang.org/x/net/websocket"
)

var fetch = utils.NewFetch()

func main() {
	go fetch.Refresh()

	s := NewServer()
	s.MountMiddleware()
	s.MountHandlers()

	utils.StartMessage(port())
	err := http.ListenAndServe(port(), s.Router)
	if err != nil {
		panic(err)
	}
}

type Server struct {
	Router *chi.Mux
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Compress(5, "text/css", "application/javascript"))

	logFormat := &utils.CustomLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags)}
	s.Router.Use(middleware.RequestLogger(logFormat))
}

func (s *Server) MountHandlers() {
	s.Router.Get("/", index)
	s.Router.Handle("/ws", websocket.Handler(wsHandler))

	fileServer := http.FileServer(http.Dir("./dist/"))
	s.Router.Handle("/public/*", http.StripPrefix("/public", fileServer))
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

func index(w http.ResponseWriter, r *http.Request) {
	// TODO cache tmpl & minify with brotli
	err := fetch.Template.Execute(w, fetch.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
}

func wsHandler(ws *websocket.Conn) {
	for {
		time.Sleep(1 * time.Minute) // Newest data already available

		err := websocket.JSON.Send(ws, fetch.Data)
		if err != nil {
			return
		}
	}
}
