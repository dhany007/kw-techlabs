package main

import (
	"fmt"
	errorpanicrecover "sesi5/error_panic_recover"
)

func main() {
	// fmt.Println("==================")
	// fmt.Println("latihan goroutine")

	// product := &product{name: "Buku", stock: 11, price: 1000}

	// go product.buy("Andi", 1)
	// go product.buy("Budi", 12)
	// go product.buy("Cahaya", 4)
	// go product.buy("Desy", 7)

	// time.Sleep(2 * time.Second)
	// fmt.Println("==================")

	// fmt.Println("==================")
	// fmt.Println("chanels implementing")
	// channels.Channels()
	// fmt.Println("\nchannel anonymous function")
	// channels.ChannelsAnynomousFunction()
	// fmt.Println("\nchannel unbuffered")
	// channels.UnbufferedChannel()
	// fmt.Println("\nchannel buffered")
	// channels.BufferedChannel()
	// fmt.Println("\nchannel select")
	// channels.ChannelSelect()
	// fmt.Println("==================")

	// fmt.Println("=============")
	// fmt.Println("defer")
	// deferexit.Defer()
	// fmt.Println("\nscope defer")
	// deferexit.ScopeDefer()
	// fmt.Println("\nexit")
	// deferexit.Exit()
	// fmt.Println("=============")

	fmt.Println("=============")
	fmt.Println("error, panic, recovery")
	errorpanicrecover.ErrorTest()
	fmt.Println()
	errorpanicrecover.CustomError()
	// fmt.Println()
	// errorpanicrecover.CustomErrorPanic()
	fmt.Println()
	errorpanicrecover.CustomErrorPanicRecover()

	fmt.Println("=============")
}

type product struct {
	name  string
	stock int
	price int
}

func (p *product) buy(name string, quantity int) {
	defer catchError()

	if quantity > p.stock {
		err := fmt.Sprintf("%s stock tidak cukup, sisa stock : %d", name, p.stock)
		panic(err)
	}

	p.stock = p.stock - quantity
	fmt.Printf("%s membeli %s sebanyak %d, total harga : %d, sisa stock: %d\n", name, p.name, quantity, quantity*p.price, p.stock)
}

func catchError() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
