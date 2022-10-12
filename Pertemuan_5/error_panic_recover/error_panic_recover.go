package errorpanicrecover

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	error : tipe data error pada go
	panic : menampilkan error dan menghentikan stack goroutine nya
	recover : untuk menangkap errornya

	jadi biasanya sepaket
*/

func ErrorTest() {
	number, err := strconv.Atoi("123m")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("number: ", number)
	}

	number, err = strconv.Atoi("123")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("number: ", number)
	}
}

func CustomError() {
	password := "andilau"

	valid, err := validatePassword(password)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func CustomErrorPanic() {
	password := "andi"

	valid, err := validatePassword(password)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(valid)
	}

}

func CustomErrorPanicRecover() {
	defer catchError()
	password := "andi"

	valid, err := validatePassword(password)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(valid)
	}

}

func validatePassword(password string) (string, error) {
	fmt.Println("password: ", password)
	if len(password) < 6 {
		return "", errors.New("password length minimal is 6")
	}

	return "password valid", nil
}

func catchError() {
	r := recover()

	if r != nil {
		fmt.Println("error occured:", r)
	} else {
		fmt.Println("apps perfectly running")
	}
}
