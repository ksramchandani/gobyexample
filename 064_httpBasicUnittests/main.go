package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", BaseHandler)
	http.ListenAndServe(":8080", nil)
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
