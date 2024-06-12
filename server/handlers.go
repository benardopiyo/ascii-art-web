package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	ascii "ascii-art-web/ascii-art"
)

var templates *template.Template

func InitTemplates() {
	var err error
	templates, err = template.ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		log.Fatal("Error parsing templates: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	style := r.FormValue("banner")

	if text == "" || style == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	art, err := ascii.GenerateArt(text, style)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Text  string
		Style string
		Art   string
	}{
		Text:  text,
		Style: style,
		Art:   art,
	}

	err = templates.ExecuteTemplate(w, "result.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
