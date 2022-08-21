package pages

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *Route {
	r := &Route{}
	r.Router = chi.NewRouter()
	r.Mount()

	return r
}

type Route struct {
	Router *chi.Mux
}

func (r *Route) Mount() {
	r.Router.Get("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	file := "./public/index.html"
	html, err := os.ReadFile(file)
	if err != nil {
		http.Error(w, "page not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
