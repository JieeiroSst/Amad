package home

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(
		"views/index.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = tmp.Execute(w,nil)
}
