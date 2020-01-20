package checkout

import (
	"html/template"
	"keikibook/src/dto"
	"log"
	"net/http"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/checkout.html")
	if r.Method != http.MethodPost {
		tmp.Execute(w, nil)
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	detail := dto.User{
		FirstName:  r.FormValue("first_name"),
		LastName:   r.FormValue("last_name"),
		CompayName: r.FormValue("company"),
		Email:      r.FormValue("email"),
		Contry:     r.FormValue("country"),
		Address:    r.FormValue("Address"),
		Town:       r.FormValue("Town"),
		ZipCode:    r.FormValue("zipCode"),
		Phone:      r.FormValue("phone"),
		Note:       r.FormValue("note"),
	}
	_ = detail
	_ = tmp.Execute(w, struct{ Success bool }{true})
}
