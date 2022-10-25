package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req, s)

	if http.StatusNotFound != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
	}

	s.MountHandlers()
	req, _ = http.NewRequest("GET", "/", nil)
	response = executeRequest(req, s)

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}
}

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}
