package latihan

import "fmt"

/*
	struct user atribut
*/

var Users []User
var Products []Product

type Shop interface {
	Save() bool
	Print()
}

type User struct {
	Name    string
	Email   string
	Address string
}

func NewUser(name, email, address string) Shop {
	return &User{
		Name:    name,
		Email:   email,
		Address: address,
	}
}

func (u *User) Save() bool {
	if u == nil {
		return false
	}

	Users = append(Users, *u)
	return true
}

func (u *User) Print() {
	for _, user := range Users {
		fmt.Println("Name: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("Address: ", user.Address)
	}
}

type Product struct {
	Name  string
	Stock int
	Price int
}

func NewProduct(name string, stock int, price int) Shop {
	return &Product{
		Name:  name,
		Stock: stock,
		Price: price,
	}
}

func (p *Product) Save() bool {
	if p == nil {
		return false
	}

	Products = append(Products, *p)
	return true
}

func (p *Product) Print() {
	for _, product := range Products {
		fmt.Println("Name: ", product.Name)
		fmt.Println("Stock: ", product.Stock)
		fmt.Println("Price: ", product.Price)
	}
}
