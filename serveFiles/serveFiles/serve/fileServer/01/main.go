package main

import (
	"io"
	"net/http"
)

func main() {
	//FileServer(http.Dir(".")) file serves from the this directory
	//will print on root page all code written on main.go
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/puppers.jpg">`)
}
