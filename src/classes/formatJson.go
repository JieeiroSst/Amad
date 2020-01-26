package classes

import (
	"encoding/json"
	"fmt"
	"keikibook/src/dto"
	"net/http"
)

var user dto.User

func FormatJsonDecoder(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "%s%s%s%s%s%s%s%s%s%s", user.FirstName, user.LastName, user.CompayName, user.Email, user.Contry, user.Address, user.Town, user.ZipCode, user.Phone, user.Note)
}

func FormatJsonEncoder(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&user)
}
