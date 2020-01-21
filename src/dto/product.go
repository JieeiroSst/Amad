package dto

type Products struct {
	Name  string
	Price string
	Note  string
	Img   []Images
}

type Images struct {
	Img string
}
