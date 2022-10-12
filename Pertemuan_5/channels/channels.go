package channels

import (
	"fmt"
	"time"
)

// channels : mekanisme untuk goroutine saling berkomunikasi dengan goroutine lainnya

func introduce(student string, c chan string) {
	result := fmt.Sprintf("my name is %s", student)

	c <- result
}

func Channels() {
	c := make(chan string)

	go introduce("dhany", c)
	go introduce("andi", c)
	go introduce("budi", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	// message yang muncul akan selalu acak
	close(c) // selalu close channel
}

// penjelasan parameter pada chan
// c chan string : menerima dan mengirim data
// c chan <- string : hanya mengirim data
// c <- chan string : hanya menerima data

func print(c <-chan string) {
	result := <-c
	fmt.Println(result)
}

func ChannelsAnynomousFunction() {
	c := make(chan string)
	students := []string{"dhany", "budi", "ani"}

	for _, student := range students {
		go func(s string) {
			fmt.Println("student:", s)
			result := fmt.Sprintf("hy my name is %s", s)
			c <- result
		}(student)
	}

	for i := 0; i < len(students); i++ {
		print(c)
	}

	close(c)
}

// contoh nya adalah yang telah kita kerjakan
// tidak mendeklarasikan banyak memori

func UnbufferedChannel() {
	c := make(chan int)

	fmt.Println("main goroutine sleep for 2 second  ")
	time.Sleep(2 * time.Second)

	go func(c chan int) {
		fmt.Println("function goroutine start sending data into channel")
		c <- 10 // disini proses pengirimannya, maka syntak dibawah ini pada fungsi go akan tertahan
		fmt.Println("function goroutine after sending data into channel")
	}(c)

	fmt.Println("main goroutine start receiving data")
	d := <-c
	fmt.Println("main goroutine received data:", d)

	close(c)
	time.Sleep(2 * time.Second) // menunggu 2 detik sebelum apps ditutup

	// print
	// main goroutine sleep for 2 second
	// main goroutine start receiving data
	// function goroutine start sending data into channel
	// main goroutine received data: 10
	// function goroutine after sending data into channel

	/*
		pada unbuffered, apapun dibawah sintak proses pengiriman data pada channel = c <- result
		akan tertahan hingga data nya diterima oleh goroutine lain
	*/
}

// menggunakan besar/jumlah memory/kapasitas yang telah dideklarasikan
func BufferedChannel() {
	c := make(chan int, 3) // akan menciptakan 3 buffer (kapasitas)

	go func(c chan int) {
		for i := 0; i < 10; i++ {
			fmt.Printf("func goroutine %d start sending data into channel\n", i)
			c <- i
			fmt.Printf("func goroutine %d after sending data into channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("main goroutine sleep for 2 seconds")
	time.Sleep(2 * time.Second)

	for v := range c { // artinya v <- c
		fmt.Println("main goroutine received value from channel", v)
	}

	/*
		proses pengiriman data akan terhenti ketika pengiriman data telah melebihi kapasitas buffer
	*/
}

func ChannelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hallo"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salut"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1: // ketika c1 sudah menerima data
			fmt.Println("Received", msg1)
		case msg2 := <-c2: // ketika c2 sudah menerima data
			fmt.Println("Received", msg2)
		}
	}
}
