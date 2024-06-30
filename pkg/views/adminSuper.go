package views

import "html/template"

func AdminSuperPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminSuper.html"))
	return t
}