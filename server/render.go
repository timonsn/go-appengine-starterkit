package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

const (
	MIMEJSON = "application/json"
	MIMEHTML = "text/html"
)

var templates *template.Template

func JSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", MIMEJSON)
	json.NewEncoder(w).Encode(obj)
}

func HTML(w http.ResponseWriter, templateName string, obj interface{}) {
	w.Header().Set("Content-Type", MIMEHTML)
	err := templates.ExecuteTemplate(w, templateName, obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func initTemplates() *template.Template {
	templates = template.New("")
	template.Must(templates.ParseGlob("templates/*/*.html"))

	return templates
}

func init() {
	initTemplates()
}
