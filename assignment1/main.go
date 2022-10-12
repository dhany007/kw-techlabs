package main

import (
	"assignment1/models"
	"assignment1/params"
	"assignment1/repositories"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Assignment 1")
	fmt.Println()

	var users []models.User

	reqDhany := params.NewCreateUser("Dhany", "dhany@koinworks.com", "Backend Engineer")
	reqAlbert := params.NewCreateUser("Albert", "albert@koinworks.com", "Backend Engineer")
	reqPutra := params.NewCreateUser("Putra", "putra@koinworks.com", "Backend Engineer")
	reqPhilip := params.NewCreateUser("Philip", "philip@koinworks.com", "Backend Engineer")

	users = append(users, repositories.InsertUser(reqDhany))
	users = append(users, repositories.InsertUser(reqAlbert))
	users = append(users, repositories.InsertUser(reqPutra))
	users = append(users, repositories.InsertUser(reqPhilip))

	// diharapkan perintah pada terminal = go run main.go 1
	if len(os.Args) < 2 {
		fmt.Println("Argument Terminal tidak valid, harap tambahkan flag nomor absen.")
		return
	}

	argument := os.Args[1]
	absen, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Argument ke 2 pada Terminal tidak valid. Harus Angka.")
		return
	}

	user, isExist := repositories.FindUserByNomorAbsen(users, absen)
	if !isExist {
		fmt.Println("User nomor absen", absen, "tidak ditemukan.")
		return
	}

	user.PrintUser()
}
