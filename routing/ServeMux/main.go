package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "woof woof")
}

type cat int

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow meow")
}

func main() {
	var d dog
	var c cat
	//newservemux gives pointer to a serve mux
	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	//listening on localhost 8080
	http.ListenAndServe(":8080", mux)
}
