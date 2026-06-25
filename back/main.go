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

func loadStatic() {
	fsStatic := http.FileServer(http.Dir("../front/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	fsPages := http.FileServer(http.Dir("../front/static"))
	http.Handle("/pages/", http.StripPrefix("/pages/", fsPages))

	fsImages := http.FileServer(http.Dir("../front/images"))
	http.Handle("/images/", http.StripPrefix("/images/", fsImages))
}

func loadSite() {
	log.Println("Сервер запущен на http://localhost:8080")
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
		Type: "Mainpage",
	}

	tmpl.Execute(w, data)
}

func main() {
	fileServer := http.FileServer(http.Dir("../front/"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			loadMainPage(w, r)
			return
		}

		fileServer.ServeHTTP(w, r)
	})

	loadStatic()
	loadSite()
}
