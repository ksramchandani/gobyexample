package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOldHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"campoy@golang.org", "old gopher campoy"},
		{"abcd", "new gopher"},
	}

	for _, c := range cases {
		request, err := http.NewRequest(
			http.MethodGet,
			"http://127.0.0.1/"+c.in,
			nil,
		)
		handleError(err, "Cannot create new request")

		recorder := httptest.NewRecorder()

		baseHandler(recorder, request)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected 200. Got %d", recorder.Code)
		}

		if !strings.Contains(recorder.Body.String(), c.out) {
			t.Errorf("Expected %s. Got %s\n", c.out, recorder.Body.String())
		}
	}
}
