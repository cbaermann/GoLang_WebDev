package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	c = context.WithValue(c, "userID", 777)
	c = context.WithValue(c, "fname", "Bond")

	results, err := dbAcess(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAcess(c context.Context) (int, error) {
	c, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := c.Value("userID").(int)
		time.Sleep(10 * time.Second)

		if c.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-c.Done():
		return 0, c.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	log.Println(c)
	fmt.Fprintln(w, c)
}
