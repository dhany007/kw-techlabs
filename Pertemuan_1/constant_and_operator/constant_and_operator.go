package constantandoperator

import "fmt"

func ConstantAndOperator() {
	fmt.Println("Constant")
	const fullName string = "Ani"
	fmt.Printf("Hello %s\n", fullName)

	fmt.Println("\nOperator")
	value := 2 + 3*4
	fmt.Println("value = ", value)

	numberOne := 1
	numberTwo := 2
	result := numberOne > numberTwo
	fmt.Println(result)

	condition := false == true
	fmt.Println(condition)

	condition = false != true
	fmt.Println(condition)

	condition = false || true
	fmt.Println(condition)
}
