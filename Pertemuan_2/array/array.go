package array

import "fmt"

func Array() {
	fmt.Println("array")

	/*
		array memiliki fixed length, dan harus diinisialisasikan terlebih dahulu panjangnya
		pakai %#v
	*/

	var numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"andi", "budi", "dhany"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

	fmt.Println("array modify through index")

	var fruits = [4]string{"apel", "nanas", "jeruk", "pepaya"}
	fmt.Printf("%#v\n", fruits)
	fruits[0] = "tomat"
	fmt.Printf("after modify = %#v", fruits)
}

func ArrayLoop() {
	fmt.Println("array loop through elements")

	var fruits = [3]string{"appel", "anggur", "tomat"}
	fmt.Println("first")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	}

	fmt.Println("second")
	for i, v := range fruits {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}
}

func ArrayMultidimensional() {
	fmt.Println("array multidimensional")

	var numbers = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, arr := range numbers {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
