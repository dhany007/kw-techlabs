package main

import (
	"fmt"
	struct_go "sesi3/struct"
)

func main() {
	// fmt.Println("latihan pertemuan 3")

	// var users []models.User

	// user1 := params.CreateUser{Name: "Dhany", Email: "dhany.aritonang@koinworks.com"}

	// users = append(users, *repositories.InsertUser(&user1))

	// for _, user := range users {
	// 	user.PrintUser()
	// }

	// fmt.Println("fungsi with return value")
	// names := []string{"dhany", "aritonang"}
	// result := function.SayHelloNameReturn("hallo", names)
	// fmt.Println(result)

	// fmt.Println("\nfungsi with multiple return value")
	// diameter := 10
	// luas, keliling := function.Calculate(float64(diameter))
	// fmt.Println(luas, keliling)

	// fmt.Println("\nfungsi with predefined return value")
	// diameter = 10
	// luas, keliling = function.Calculate(float64(diameter))
	// fmt.Println(luas, keliling)

	// fmt.Println("\nfungsi with  variadic parameter")
	// students := function.StudentsVariadic("andi", "budi", "ani")
	// fmt.Printf("%v", students)

	// fmt.Println("\n\nfungsi with  variadic parameter slice")
	// numbers := []int{1, 2, 3, 4, 5}
	// total := function.SumVariadic(numbers...)
	// fmt.Println("total =", total)

	// fmt.Println("\nfungsi with  variadic parameter slice")
	// name := "dhany"
	// favFoods := []string{"mie", "nasi", "ayam"}
	// function.Profile(name, favFoods...)

	// fmt.Println("closure")
	// fmt.Println("\ndeclare closure")
	// closure.DeclareClossure()

	// fmt.Println("\nclosure IIFE => langsung dijalankan")
	// closure.ClosureIIFE()

	// fmt.Println("\nclosure as a return function")
	// students := []string{"dhanY", "ANI", "budi"}
	// findValue := closure.FindStudent(students)
	// fmt.Println(findValue("DHANY"))

	// fmt.Println("\nclosure as paramater or callback")
	// numbers := []int{1, 2, 3, 4, 7, 8}
	// findOdd := func(number int) bool {
	// 	return number%2 == 1
	// }
	// fmt.Println(closure.FindOddNumber(numbers, findOdd))

	// fmt.Println("\nclosure as paramater or callback alias")
	// numbers = []int{1, 2, 3, 4, 7, 8}
	// findOdd = func(number int) bool {
	// 	return number%2 == 1
	// }
	// fmt.Println(closure.FindOddNumberAlias(numbers, findOdd))

	// fmt.Println("pointer")
	// fmt.Println("\npointer memory address")
	// pointer.PointerMemoryAddress()

	// fmt.Println("\npointer as parameter")
	// a := 10
	// fmt.Println("Before:", a)
	// pointer.PointerAsParameter(&a)
	// fmt.Println("After:", a)

	fmt.Println("struct")
	fmt.Println("\ngiving value struct")
	struct_go.GivingValueStruct()

	fmt.Println("\ninitializing struct")
	struct_go.InitializingStruct()

	fmt.Println("\npointer to struct")
	struct_go.PointerToStruct()

	fmt.Println("\nembedded struct")
	struct_go.EmbeddedStruct()

	fmt.Println("\nAnynomous struct")
	struct_go.AnynomousStruct()

	fmt.Println("\nslice of struct")
	struct_go.SliceOfStruct()

	fmt.Println("\nslice of anoonymous struct")
	struct_go.SliceAnonymousStruct()
}
