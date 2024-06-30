package views

import "html/template"

func ClientHistoryPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/clientHistory.html"))
	return t
}