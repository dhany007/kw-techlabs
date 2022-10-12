package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http request")

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/a")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Body)

}

// func httpCall() {

// }
