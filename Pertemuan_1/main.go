package main

import (
	"fmt"
	constantandoperator "sesi1/constant_and_operator"
	"sesi1/datatypes"
	"sesi1/test"
	"sesi1/variabel"
)

func main() {
	fmt.Println("Sesi 1")

	fmt.Println("Hello World")
	test.Test()

	fmt.Println("\nVariabel")
	variabel.VariabelWithDataTypes()
	variabel.VariabelWithoutDataTypes()
	variabel.MultipleVariabel()
	variabel.UnderscoreVariabel()
	variabel.VariabelPrintF()

	fmt.Println("\n\nData Types")
	datatypes.DataTypesBasic()

	fmt.Println("\nConstant and Operator")
	constantandoperator.ConstantAndOperator()
}
