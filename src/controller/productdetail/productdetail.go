package productdetail

import (
	"html/template"
	"log"
	"net/http"
)

func ProductDetail(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/product-details.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmp.Execute(w,nil)
}
