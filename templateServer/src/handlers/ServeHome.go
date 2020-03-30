package handlers

import (
	"net/http"
	"time"
	m "../models"
	"html/template"
	"path"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(path.Dir("../assets/test.html")))

	now := time.Now()
	homeTemplateVars := m.TemplateModel{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	tmpl.Execute(w, homeTemplateVars)
}