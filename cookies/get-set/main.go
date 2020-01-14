package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "Cookie Written - Check Your Browser")
	fmt.Fprintln(w, "in chrome go to : dev tools/application/cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie #1: ", c1)
	}
	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie #2: ", c2)
	}
	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie #3: ", c3)
	}
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})
	fmt.Fprintln(w, "Cookies Written - Check Your Browser")
	fmt.Fprintln(w, "in chrome go to: dev tools/application/cookies")
}
