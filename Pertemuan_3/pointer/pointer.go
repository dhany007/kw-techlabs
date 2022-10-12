package pointer

import "fmt"

func PointerMemoryAddress() {
	firstNumber := 4
	secondNumber := &firstNumber

	fmt.Println("firstNumber value", firstNumber)
	fmt.Println("fitstNumber memory address", &firstNumber)

	fmt.Println("secondNumber value", *secondNumber)
	fmt.Println("secondNumber memory address", secondNumber)
}

func PointerAsParameter(number *int) {
	*number = 20
}
