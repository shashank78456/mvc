package views

import "html/template"

func ClientHistoryPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/clientHistory.html"))
	return t
}

func ClientHomePage() *template.Template {
	t := template.Must(template.ParseFiles("templates/clientHome.html"))
	return t
}

func ClientReturnPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/clientReturn.html"))
	return t
}
