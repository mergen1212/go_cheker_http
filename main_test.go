package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handelRequest(w *httptest.ResponseRecorder, r *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, r)
}
func TestNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/badrequest", nil)
	recorder := httptest.NewRecorder()
	handelRequest(recorder, request)
	if recorder.Code != 404 {
		t.Error("Must be 404 but", recorder.Code)
	}
}
