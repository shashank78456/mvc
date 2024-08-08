package views

import (
	"html/template"
)

func LoginPage() *template.Template {
	t := template.Must(template.ParseFiles(("templates/login.html")))
	return t
}

func SignupPage() *template.Template {
	t := template.Must(template.ParseFiles(("templates/signup.html")))
	return t
}
