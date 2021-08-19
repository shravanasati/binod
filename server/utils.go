package main

import (
	"html/template"
	"net/http"
)

type indexPage struct {
	stylePath string
	scriptPath string
}

func renderTemplate(filename string, w http.ResponseWriter, obj interface{}) {
	t := template.Must(template.ParseFiles(filename))
	t.Execute(w, obj)
}