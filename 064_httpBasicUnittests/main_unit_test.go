package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test001_BasicCase(t *testing.T) {

	request, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("Cannot create http NewRequest")
	}

	response := httptest.NewRecorder()
	BaseHandler(response, request)

	want := "hello"
	got := response.Body.String()

	if got != want {
		t.Errorf("Want %s Got %s", want, got)
	}

}
