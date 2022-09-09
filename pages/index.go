package pages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var fetch = NewFetch()
var tmpl, _ = LoadTemplate()

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
}

func index(w http.ResponseWriter, r *http.Request) {
	if fetch.Refresh() {
		fetch.Update()
	}

	err := tmpl.Execute(w, fetch.data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
}
