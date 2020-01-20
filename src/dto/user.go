package dto

type User struct {
	FirstName  string `json:firstname`
	LastName   string `json:lastname`
	CompayName string `json:compayname`
	Email      string `json:email`
	Contry     string `json:contry`
	Address    string `json:address`
	Town       string `json:town`
	ZipCode    string `json:zipCode`
	Phone      string `json:phone`
	Note       string `json:note`
}
