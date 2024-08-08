package views

import (
	"html/template"
)

func AdminHomePage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminHome.html"))
	return t
}

func AdminPromptPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminPrompt.html"))
	return t
}

func AdminRequestsPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminRequests.html"))
	return t
}
