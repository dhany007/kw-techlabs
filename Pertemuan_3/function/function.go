package function

import (
	"fmt"
	"math"
	"strings"
)

func SayHello() {
	fmt.Println("Hallo")
}

func SayHelloName(name string, age string) {
	fmt.Println("Hello ", name, " age: ", age)
}

func SayHelloNameReturn(msg string, names []string) string {
	joinStr := strings.Join(names, " ")

	// sprintf => mereturn sebuah nilai
	result := fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

func Calculate(diameter float64) (float64, float64) {
	area := math.Pi * math.Pow(diameter/2, 2) // luas = phi * r pangkat 2
	circumFerence := math.Pi * diameter       // keliling = 2 phi r

	return area, circumFerence
}

func CalculatePredefinedValue(diameter float64) (area float64, circumFerence float64) {
	area = math.Pi * math.Pow(diameter/2, 2) // luas = phi * r pangkat 2
	circumFerence = math.Pi * diameter       // keliling = 2 phi r

	return
}

// hanya satu parameter yang diperbolehkan mengggunakan variadic dan harus diletakkan diakhir
// return array of map (array of object)
func StudentsVariadic(names ...string) []map[string]string {
	students := []map[string]string{}

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}

		students = append(students, temp)
	}

	return students
}

func SumVariadic(number ...int) int {
	total := 0

	for _, v := range number {
		total += v
	}

	return total
}

// hanya satu parameter yang diperbolehkan mengggunakan variadic dan harus diletakkan diakhir
func Profile(name string, foods ...string) {
	mergeFavFoods := strings.Join(foods, ", ")

	fmt.Println("hello, my name is ", name)
	fmt.Println("my favorite foods are", mergeFavFoods)
}
