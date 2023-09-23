package handlers

import (
	"html/template"
	"net/http"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl", "templates/home.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
