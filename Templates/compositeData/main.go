package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl3.gohtml"))
}

func main() {
	// sages := []string{"Ghandi", "MLK", "Buddha", "Jesus", "Muhammed"}

	sages := map[string]string{
		"India":    "Ghandi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhamad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
