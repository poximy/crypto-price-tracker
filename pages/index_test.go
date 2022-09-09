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
	if val, ok := headers["Content-Type"]; ok {
		if len(val)	!= 1 {
			t.Errorf("Expected Content-Type Header to be of length 1. Got %d\n", len(val))
		}

		contains := []string{"text/html"}
		equal := reflect.DeepEqual(val, contains)
		if !equal {
			t.Errorf("Expected %+q to be the same as %+q", val, contains)
		}
	}

}

func executeRequest(req *http.Request, r *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}
