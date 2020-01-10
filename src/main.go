package main

import (
	"keikibook/src/controller/cart"
	"keikibook/src/controller/checkout"
	"keikibook/src/controller/home"
	"keikibook/src/controller/productdetail"
	"keikibook/src/controller/shop"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/Amad", home.Index)
	http.HandleFunc("/Amad/cart", cart.Cart)
	http.HandleFunc("/Amad/checkout", checkout.CheckOut)
	http.HandleFunc("/Amad/productdetail", productdetail.ProductDetail)
	http.HandleFunc("/Amad/shop", shop.Shop)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
