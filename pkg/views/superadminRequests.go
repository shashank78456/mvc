package views

import (
	"html/template"
)

func SuperAdminRequestsPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminRequests.html"))
	return t
}