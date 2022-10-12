package aliases

import "fmt"

// alias digunakan untuk alias dari suatu datatypes
type configKey string
type constants string

func Aliases() {
	var DB_HOST configKey = "localhost"
	var DB_PORT constants = "5432"

	fmt.Println(DB_HOST)
	fmt.Println(DB_PORT)
}
