package main

import (
	"io"
	"net/http"
)

//after launching docker container
//open on localhost, not localhost:80
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from a docker container")
}
