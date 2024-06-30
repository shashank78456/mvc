package views

import (
	"html/template"
)

func ClientHomePage() *template.Template {
	t := template.Must(template.ParseFiles("templates/clientHome.html"))
	return t
}