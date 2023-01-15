package pages

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/net/websocket"
)

var fetch = NewFetch()

func NewRoute() *chi.Mux {
	r := &Route{}
	r.Router = chi.NewRouter()
	r.Mount()

	return r.Router
}

type Route struct {
	Router *chi.Mux
}

func (r *Route) Mount() {
	r.Router.Get("/", index)
	r.Router.Handle("/ws", websocket.Handler(wsHandler))

	fileServer := http.FileServer(http.Dir("./public/"))
	r.Router.Handle("/public/*", http.StripPrefix("/public", fileServer))
}

func index(w http.ResponseWriter, r *http.Request) {
	fetch.Refresh()

	// TODO cache tmpl & minify
	err := fetch.template.Execute(w, fetch.data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
}

func wsHandler(ws *websocket.Conn) {
	for {
		time.Sleep(30 * time.Second)
		fetch.Refresh()

		err := websocket.JSON.Send(ws, fetch.data)
		if err != nil {
			return
		}
	}
}
