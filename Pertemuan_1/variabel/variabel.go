package variabel

import "fmt"

func VariabelWithDataTypes() {
	fmt.Println("\nVariabel With Data Types")

	var name string = "dhany"
	var age int = 2

	fmt.Println("ini adalah namanya ==>", name)
	fmt.Println("ini adalah umurnya ==>", age)

	var name2 string
	var age2 int = 24

	name2 = "budi"
	age2 = 23

	fmt.Println("ini adalah namanya ==>", name2)
	fmt.Println("ini adalah umurnya ==>", age2)

}

func VariabelWithoutDataTypes() {
	fmt.Println("\nVariabel Without Data Types")

	var name string = "dhany"
	var age int = 2

	fmt.Println("ini adalah namanya ==>", name)
	fmt.Println("ini adalah umurnya ==>", age)

	var name2 string
	var age2 int = 24

	name2 = "budi"
	age2 = 23

	fmt.Printf("%T %T", name2, age2)

	var name3 = "ani"
	var age3 = 23

	fmt.Printf("\n%T %T", name3, age3)

	fmt.Println("\nShort Declaration")

	name4 := "kalai"
	age4 := 20
	fmt.Printf("%T %T", name4, age4)
}

func MultipleVariabel() {
	fmt.Println("\nMultiple Variable")

	var name, age, address = "dhany", 23, "jakarta"
	first, second, third := 1, 2, 3

	fmt.Println(name, age, address)
	fmt.Println(first, second, third)
}

func UnderscoreVariabel() {
	fmt.Println("\nUnderscore Variable")

	_, _, name := "tidak", "tidak", "ya"

	fmt.Println("digunakan ", name)
	fmt.Println("digunakan untuk mendeklarasikan variabel yang tidak digunakan")
}

func VariabelPrintF() {
	fmt.Println("\nPrintF")

	name, age, address := "Dhany", 23, "Jakarta"

	fmt.Printf("nama %s, age %d, address %s", name, age, address)

}
