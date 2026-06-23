package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
)

type PageData struct {
	Name string
	Type string
}

func loadCSSandJS() {
	fs := http.FileServer(http.Dir("../front/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
}

func loadSite() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func loadMainPage(w http.ResponseWriter, r *http.Request) {
	mainHtmlFile := "../front/index.html"
	if _, err := os.Stat(mainHtmlFile); os.IsNotExist(err) {
		log.Fatal("файл: ", mainHtmlFile, "не найден")
	}
	
	tmpl, err := template.ParseFiles(mainHtmlFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := PageData { 
		Name: "OpenDelta",
		Type: "MainPage",		
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", loadMainPage)
	loadCSSandJS()
	loadSite()	
}
