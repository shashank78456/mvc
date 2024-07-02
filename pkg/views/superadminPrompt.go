package views

import (
	"html/template"
)

func SuperAdminPromptPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/superadminPrompt.html"))
	return t
}