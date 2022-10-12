package main

import (
	"fmt"
	"sesi4/latihan"
	"sync"
)

func main() {
	// fmt.Println("interface")
	// fmt.Println()

	// var rectangle interface4.Shape = interface4.Rectangle{Width: 10.5, Height: 20.5}
	// var circle interface4.Shape = interface4.Circle{Radius: 10.5}

	// r := interface4.NewRectangle(10, 20)
	// c := interface4.NewCircle(10)
	// fmt.Println("r:", r.Area())
	// fmt.Println("c:", c.Area())
	// fmt.Printf("rectangle: %T\n", rectangle)
	// fmt.Printf("circle: %T\n", circle)

	// interface4.Print("circle", circle)
	// interface4.Print("rectangle", rectangle)

	// fmt.Println("\nassertion interface")

	// // circle merupakan interface shape
	// // ketika kita memanggil volume, akan error
	// // untuk mencegah error, maka dibuat assertion, yaitu dipaksa ke tipe data aslinya

	// // circle.(interface4.Circle).Volume()
	// // dapat dicek apakah berhasil atau tidak
	// value, ok := circle.(interface4.Circle)
	// if ok {
	// 	fmt.Println("circle value:", value)
	// 	fmt.Println("circle volume:", value.Volume())
	// }

	// fmt.Println("empty interface")
	// emptyinterface.EmptyInterface()

	// fmt.Println("\nempty interface assertion")
	// emptyinterface.EmptyInterfaceAssertion()

	// fmt.Println("\nmap and slice empty interface")
	// emptyinterface.EmptyInterfaceMapSlice()

	// fmt.Println("reflect")
	// fmt.Println("\nidentify data type and value")
	// reflect4.IdentifyDataTypeAndValue()

	// fmt.Println("\nidentify method")
	// reflect4.IdentyfingMethod()
	// products := []interface{}{}

	fmt.Println("==================")

	fmt.Println("latihan interface")

	users := []latihan.Shop{
		latihan.NewUser("dhany", "dhany@koinworks.com", "jkt"),
		latihan.NewUser("adni", "andi@koinworks.com", "jkt"),
		latihan.NewUser("cantika", "cantika@koinworks.com", "jkt"),
	}

	for _, user := range users {
		user.Save()
	}

	if len(users) > 0 {
		fmt.Println("users :")
		users[0].Print()
	}

	products := []latihan.Shop{
		latihan.NewProduct("pensil", 10, 100000),
		latihan.NewProduct("buku", 5, 200000),
		latihan.NewProduct("penghapus", 5, 200000),
	}

	for _, product := range products {
		product.Save()
	}

	if len(products) > 0 {
		fmt.Println("\nproduct :")
		products[0].Print()
	}

	fmt.Println("==================")

	// fmt.Println("=========")

	// fmt.Println("waitgroup goroutine")
	// fmt.Println()

	// fruits := []string{"apel", "jeruk", "tomat", "anggur"}
	// wg2 := sync.WaitGroup{} // mendeklarasikan waitgroup

	// for i, v := range fruits {
	// 	wg2.Add(1)
	// 	go printFruit(i, v, &wg2)
	// }
	// wg2.Wait()

	// fmt.Println("=========")

	// fmt.Println("\nlatihan goroutine")
	// fmt.Println("\n main execution started")
	// wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go func() {
	// 	firstProcess(8)
	// 	wg.Done()
	// }()

	// wg.Wait()
	// secondProcess(9)

	// fmt.Println("\n main execution ended")
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index: %d, fruit: %s \n", index, fruit)
	wg.Done()
}

func firstProcess(index int) {
	fmt.Println("first process start ")

	for i := 0; i < index; i++ {
		fmt.Println("i =", i)
	}
	fmt.Println("first process done")
}

func secondProcess(index int) {
	fmt.Println("second process start ")
	for j := 0; j < index; j++ {
		fmt.Println("j =", j)
	}
	fmt.Println("second process end")
}
