package views

import (
	"html/template"
)

func SignupPage() *template.Template {
	t := template.Must(template.ParseFiles(("templates/signup.html")))
	return t
}