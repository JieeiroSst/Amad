package shop

import (
	"html/template"
	"log"
	"net/http"
)

func Shop(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/shop.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmp.Execute(w,nil)
}
