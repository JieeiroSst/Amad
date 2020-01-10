package checkout

import (
	"html/template"
	"log"
	"net/http"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/checkout.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmp.Execute(w,nil)
}
