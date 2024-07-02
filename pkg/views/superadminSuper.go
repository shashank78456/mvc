package views

import "html/template"

func SuperAdminSuperPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminSuper.html"))
	return t
}