package slice

import "fmt"

func Slice() {
	fmt.Println("slice")
	// tidak memiliki fixed-length
	// reference type, kalau dicopy, dan copyannya diubah, maka aslinya juga ikut berubah

	var fruits = []string{"tomat", "cabe", "apel"}
	fmt.Printf("%#v\n", fruits)

	fmt.Println("slice make function")
	var fruits2 = make([]string, 3)
	fmt.Printf("%#v\n", fruits2)

	fmt.Println("slice append function")
	var fruits3 = make([]string, 3)
	fmt.Printf("%#v\n", fruits3)
	fruits3 = append(fruits3, "bunga", "indah", "sari")
	fmt.Printf("fruits setelah diappend : %#v\n", fruits3)

	var fruits5 = []string{"kol", "sawi", "toge"}
	fmt.Printf("fruits5: %#v\n", fruits5)
	var fruits6 = []string{"bayam"}
	fmt.Printf("fruits6: %#v\n", fruits6)
	fruits5 = append(fruits5, fruits6...)
	fmt.Printf("fruits5 setelah diapend dengan fruits6: %#v\n", fruits5)

	fmt.Println("copy function")
	var foods = []string{"sambal", "daging"}
	var drinks = []string{"jus", "air putih", "susu"}
	fmt.Printf("foods : %#v\n", foods)
	fmt.Printf("drinks : %#v\n", drinks)
	nn := copy(foods, drinks)
	fmt.Printf("foods : %#v\n", foods)
	fmt.Printf("drinks : %#v\n", drinks)
	fmt.Println("n : ", nn)

}
