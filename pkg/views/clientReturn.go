package views

import "html/template"

func ClientReturnPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/clientReturn.html"))
	return t
}