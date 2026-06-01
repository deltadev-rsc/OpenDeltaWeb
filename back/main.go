package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
)

type PageData struct {
	Name string
}

func renderHtml(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("../front/index.html"); os.IsNotExist(err) {
		log.Fatal("файл index.html не найден")
	}
	
	tmpl, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
