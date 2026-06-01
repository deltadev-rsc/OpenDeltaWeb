package main

import (
	"net/http"
	"html/template"
)

type PageData struct {
	Name string
}

func renderHtml(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData { Name: "OpenDelta" }
	tmpl.Execute(w, data)
}

func backStart() {
	http.HandleFunc("/", renderHtml)
	http.ListenAndServe(":8080", nil)
}

func main() {
	backStart()
}
