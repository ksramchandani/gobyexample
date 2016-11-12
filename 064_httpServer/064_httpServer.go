/*
Usage: go run 064_httpServer.go to start http server
In another terminal window: curl 127.0.0.1:8080/

*/

package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func handleError(e error, msg string) {
	if e != nil {
		fmt.Printf("%s : %s", msg, e)
	}
}

func main() {
	http.HandleFunc("/", baseHandler)
	err := http.ListenAndServe(":8080", nil)
	handleError(err, "Cannot start HTTP listener")

}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)
	if match != nil {
		fmt.Fprintf(w, "Welcome old gopher %s\n", match[0][1])

	} else {
		fmt.Fprintln(w, "Welcome new gopher")
	}
}
