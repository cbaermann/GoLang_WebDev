package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@test.com")
	fmt.Println(c)
	c = getCode("test.test1.com")
	fmt.Println(c)
}

//pass in string, return of type string
func getCode(s string) string {
	//gives back hash
	h := hmac.New(sha256.New, []byte("ourkey"))
	//write string into the hash
	io.WriteString(h, s)
	//string print to hexadecimal
	return fmt.Sprintf("%x", h.Sum(nil))
}
