package stringsindepth

import "fmt"

func Strings() {
	fmt.Println("strings in depth")

	city := "jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Println("index: ", i, ", byte: ", city[i])
	}

	var cityByte []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println("strings: ", string(cityByte))
}
