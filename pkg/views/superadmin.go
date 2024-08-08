package views

import (
	"html/template"
)

func SuperAdminHomePage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminHome.html"))
	return t
}

func SuperAdminPromptPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminPrompt.html"))
	return t
}

func SuperAdminRequestsPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminRequests.html"))
	return t
}

func SuperAdminSuperPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminSuper.html"))
	return t
}
