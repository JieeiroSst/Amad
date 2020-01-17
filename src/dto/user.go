package dto

type User struct {
	firstName  string `json:firstname`
	lastName   string `json:lastname`
	compayName string `json:compayname`
	email      string `json:email`
	contry     string `json:contry`
	address    string `json:address`
	town       string `json:town`
	zipCode    string `json:zipCode`
	phone      string `json:phone`
	note       string `json:note`
}
