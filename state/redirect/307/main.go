package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at foo: ", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request method at bar: ", r.Method)
	//process form submission here
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred: ", r.Method)
	//redirects to html page through form on index.html
	// which takes action of ="/bar"
	//renders index.html
	tpl.ExecuteTemplate(w, "index.html", nil)
}
