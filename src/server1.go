// go run ./src/server1.go &  // to start the server in background
// ps -l // to check running processes
// view in localhost:8000 in browser, or use
//./bin/fetch localhost:8000/testing1234

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls a handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handles echoes the path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
