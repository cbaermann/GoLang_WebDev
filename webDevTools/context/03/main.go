package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c, cancel := context.WithCancel(context.Background())
	defer cancel() //make sure all paths cancel context to avoid leak

	for n := range gen(c) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

func gen(c context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-c.Done():
				return //avoid leaking of goroutine when c is finished
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
