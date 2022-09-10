package pages

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNewRoute(t *testing.T) {
	r := Route{Router: chi.NewRouter()}

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req, r.Router)

	if http.StatusNotFound != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
	}

	r.Mount()
	req, _ = http.NewRequest("GET", "/", nil)
	response = executeRequest(req, r.Router)

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
	}
}

func TestIndex(t *testing.T) {
	r := NewRoute()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req, r)

	headers := response.Header()
	contentType, ok := headers["Content-Type"]
	if !ok {
		t.Error("Expected Content-Type in Headers. Content-Type not found\n")
		return
	}

	if length := len(contentType); length	!= 1 {
		t.Errorf("Expected Content-Type Header to be of length 1. Got %d\n", length)
	}

	contains := []string{"text/html"}
	equal := reflect.DeepEqual(contentType, contains)
	if !equal {
		t.Errorf("Expected %+q to be the same as %+q\n", contentType, contains)
	}
}

func executeRequest(req *http.Request, r *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}
