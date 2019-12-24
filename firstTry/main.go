package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	//creates index.html file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Panicln("error creating file")
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

//used go run main.go > index.html to transform html in main.go to it's own html file
