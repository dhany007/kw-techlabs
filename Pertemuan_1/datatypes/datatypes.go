package datatypes

import "fmt"

func DataTypesBasic() {
	fmt.Println("datatypes basic")
	fmt.Println("\nNumber (Integer)")

	var first = 20
	var second = -20

	fmt.Printf("\ntipe data first = %T", first)
	fmt.Printf("\ntipe data second = %T", second)

	var first2 uint = 10
	var second2 int = -20
	fmt.Printf("\ntipe data first = %T", first2)
	fmt.Printf("\ntipe data second = %T\n", second2)

	fmt.Println("\nNumber (Floating)")
	var decimalNumber float32 = 3.14
	fmt.Printf("decimal number : %f \n", decimalNumber)
	fmt.Printf("decimal number : %.3f \n", decimalNumber)

	fmt.Println("\nBoolean")
	var isPermitted bool = false
	fmt.Printf("is it permitted ? %t \n", isPermitted)

}
