package main

import (
	"fmt"
	"latihan/controllers"
	"net/http"
)

func main() {
	fmt.Println("latihan sesi 9")

	PORT := ":3000"

	http.HandleFunc("/users", controllers.GetAllUsers)
	http.HandleFunc("/users/register", controllers.Register)

	fmt.Println("server running on port", PORT)

	http.ListenAndServe(PORT, nil)
}
