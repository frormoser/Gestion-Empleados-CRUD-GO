package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Start)

	log.Println("Server running...")

	http.ListenAndServe(":8080", nil)
}
func Start(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "home", nil)
}
