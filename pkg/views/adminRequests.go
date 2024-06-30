package views

import (
	"html/template"
)

func AdminRequestsPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminRequests.html"))
	return t
}