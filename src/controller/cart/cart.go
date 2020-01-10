package cart

import (
	"html/template"
	"log"
	"net/http"
)

func Cart(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/cart.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmp.Execute(w,nil)
}
