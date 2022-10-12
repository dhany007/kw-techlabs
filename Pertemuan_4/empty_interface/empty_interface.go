package emptyinterface

import "fmt"

func EmptyInterface() {
	// suatu tipe data yang dapat menerima semua tipe data
	var randomValues interface{}

	_ = randomValues

	randomValues = "20"
	randomValues = 20
	randomValues = []string{"andi"}
}

func EmptyInterfaceAssertion() {
	// memaksa empty interface ke suatu tipe data
	var v interface{}

	v = 20

	// memaksa ke int
	value, ok := v.(int)

	if ok {
		v = value * 20
	}

	fmt.Println("v :", v)
}

func EmptyInterfaceMapSlice() {
	rs := []interface{}{"andi", "budi"}

	rm := map[string]interface{}{
		"Nama": "dhany",
		"age":  10,
	}

	fmt.Println("rs :", rs)
	fmt.Println("rm :", rm)
}
