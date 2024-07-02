package views

import (
	"html/template"
)

func SuperAdminHomePage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminHome.html"))
	return t
}