package models

var Products []Product

type Product struct {
	Nama  string `json:"nama"`
	Brand string `json:"brand"`
	Stok  int    `json:"stok"`
	Price int    `json:"price"`
}

func AddProduct(u *User) {
	Users = append(Users, *u)
}

func GetProducts() *[]Product {
	return &Products
}
