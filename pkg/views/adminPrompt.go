package views

import (
	"html/template"
)

func AdminPromptPage() *template.Template {
	t := template.Must(template.ParseFiles("templates/adminPrompt.html"))
	return t
}