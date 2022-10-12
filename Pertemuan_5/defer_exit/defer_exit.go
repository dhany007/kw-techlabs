package deferexit

import (
	"fmt"
	"os"
)

/*
	defer : dijalankan diakhir scope fungsinya
	exit : memaksa menghentikan programnya, dan sintak dibawah nya tidak akan dibaca
*/
func Defer() {
	defer fmt.Println("defer function start to execute") // ini akan dijalankan ketika baris fungsi nya selesai
	fmt.Println("hy everyone")
	fmt.Println("welcome back")

	// hasil
	// hy everyone
	// welcome back
	// defer function start to execute
}

func ScopeDefer() {
	callDeferFunc()
	fmt.Println("main scope")

	/*
		hasil

		scope defer function
		defer function executed
		main scope

		defer nya ada difungsi lain, jadi tidak terpengaruh pada main
	*/
}

func callDeferFunc() {
	defer deferFunction()
	fmt.Println("scope defer function")
}

func deferFunction() {
	fmt.Println("defer function executed")
}

func Exit() {
	defer fmt.Println("invoke defer")
	fmt.Println("hallo")
	os.Exit(1)
	// fungsi defer diatas tidak akan dijalankan
	// karena fungsi defer seharusnya dijalankan diakhir, yaitu setelah os.exit
	// tapi os.exit memaksa aplikasi nya berhenti

	// hasil
	// hallo
	// exit status 1
}
