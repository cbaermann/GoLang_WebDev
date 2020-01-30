package config

//reads html from templates folder
import "html/template"

var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.html"))
}
