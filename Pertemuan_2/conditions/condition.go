package conditions

import (
	"fmt"
)

func ConditionTemporaryVariabel() {
	fmt.Println("Kondisi Temporary Variabel")

	currentYear := 2022

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
}

func Switch() {
	fmt.Println("Switch")

	score := 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("not bad")
	}
}

func SwitchRelationalOperator() {
	fmt.Println("Switch with relational operator")

	score := 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}
}

func SwitchFallthrough() {
	fmt.Println("Switch fallthrough keyword")
	// fallthrough => melanjutkan case selanjutnya walaupun case nya sudah terpenuhi

	score := 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Awesome")
		fallthrough
	case score < 5:
		fmt.Println("it is ok, but please study harder")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You don't have a good score yet.")
		}
	}
}

func NestedConditions() {
	fmt.Println("Nested conditions")

	score := 8

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}

func LatihanKondisi() {
	age := 28
	gender := "female"

	canWork := false

	if gender == "male" {
		switch {
		case age > 17:
			canWork = true
		default:
			canWork = false

		}
	} else {
		switch {
		case age > 20:
			canWork = true
		default:
			canWork = false
		}
	}

	if canWork {
		fmt.Printf("%s berumur %d, dapat bekerja \n", gender, age)
	} else {
		fmt.Printf("%s berumur %d, belum dapat bekerja \n", gender, age)
	}
}

func LatihanKondisiGanjilGenap() {
	fmt.Println("latihan pengecekan kondisi ganjil atau genap")

	number := 10
	if number%2 == 0 {
		fmt.Println(number, "adalah bilangan Genap")
	} else {
		fmt.Println(number, "adalah bilangan Ganjil")
	}
}
