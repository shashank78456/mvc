package views

import (
	"html/template"
)

func AdminHomePage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminHome.html"))
	return t
}