package closure

import (
	"fmt"
	"strings"
)

// closure adalah function tanpa nama yang disimpan dalam variabel

func DeclareClossure() {
	evenNumber := func(number ...int) []int {
		result := []int{}
		for _, v := range number {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(evenNumber(numbers...))
}

func ClosureIIFE() {
	isPalindrome := func(str string) bool {
		temp := ""

		for i := len(str) - 1; i >= 0; i-- {
			// mengambil byte huruf perhuruf dan mengubahnya ke string
			temp += string(byte(str[i]))
		}

		return temp == str
	}("bukan") // langsung dijalankan : false, coba "katak" : true

	fmt.Println(isPalindrome)
}

// closure as return function
func FindStudent(students []string) func(string) string {
	return func(name string) string {
		student := ""
		position := 0

		for i, v := range students {
			if strings.EqualFold(name, v) {
				student = v
				position = i
			}
			break
		}

		if student == "" {
			return fmt.Sprintf("student with name %s not found", name)
		} else {
			return fmt.Sprintf("student with name %s found in position %d", name, position)
		}
	}
}

// closure as paramater fungsi or callback

// menjumlahkan bilangan ganjil, tetapi logika ganjilnya disimpan dalam fungsi callback yg mereturn bool
func FindOddNumber(numbers []int, callback func(int) bool) int {
	totalOddNumber := 0

	for _, number := range numbers {
		if callback(number) {
			totalOddNumber += number
		}
	}

	return totalOddNumber
}

// callback ditulis dalam bentuk alias

type isOddNumber func(int) bool

func FindOddNumberAlias(numbers []int, callback isOddNumber) int {
	totalOddNumber := 0

	for _, number := range numbers {
		if callback(number) {
			totalOddNumber += number
		}
	}

	return totalOddNumber
}
