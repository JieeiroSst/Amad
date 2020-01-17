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

	http.HandleFunc("/Amado", home.Index)
	http.HandleFunc("/Amado/cart", cart.Cart)
	http.HandleFunc("/Amado/checkout", checkout.CheckOut)
	http.HandleFunc("/Amado/productdetail", productdetail.ProductDetail)
	http.HandleFunc("/Amado/shop", shop.Shop)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
