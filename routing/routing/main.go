package main

import (
	"io"
	"net/http"
)

type num int

func (n num) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "bark bark bark")
	case "/cat":
		io.WriteString(w, "meow meow meow")
	}
}

func main() {
	var n num
	http.ListenAndServe(":8080", n)
}
