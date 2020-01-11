package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/puppers.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/puppers.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("puppers.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	//stat gives pointer to a file
	//gives back file info and an error
	//file info gives acess to ModTime()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
