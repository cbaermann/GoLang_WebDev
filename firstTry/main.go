package main

import "fmt"

func main() {
	name := "Todd McLeod"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}

//used go run main.go > index.html to transform html in main.go to it's own html file
