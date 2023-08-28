package web

import (
	"html/template"
	"log"
	"net/http"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html"))

	if err := temp.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
