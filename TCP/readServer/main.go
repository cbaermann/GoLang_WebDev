package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//takes a reader
	scanner := bufio.NewScanner(conn)
	//returns bool
	for scanner.Scan() {
		//gives line
		ln := scanner.Text()
		//prints line
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}
