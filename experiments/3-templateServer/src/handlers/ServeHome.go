package handlers

import (
	m "../models"
	"html/template"
	"net/http"
	"time"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	// change this
	tmpl := template.Must(template.ParseFiles(`/Users/adamgerard/Desktop/Workspace/_code/go_refresh/experiments/3-templateServer/src/assets/test.html`))

	now := time.Now()
	homeTemplateVars := m.TemplateModel{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	tmpl.Execute(w, homeTemplateVars)
}