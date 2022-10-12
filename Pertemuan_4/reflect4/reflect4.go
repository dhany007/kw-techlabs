package reflect4

import (
	"fmt"
	"reflect"
)

// reflect : melakukan inspeksi pada variabel, mengambil informasi dari variabel tersebut,
// bahkan memanipulasi nya

func IdentifyDataTypeAndValue() {
	number := 20
	reflectValue := reflect.ValueOf(number) // mengambil ke variabel reflect

	fmt.Println("type variabel: ", reflectValue.Type()) // type variabel

	// kind = mememeriksa terlebih dahulu apa jenis tipe datanya
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel: ", reflectValue.Int()) // mengambil value int
	}

	text := "andi"
	reflectValue = reflect.ValueOf(text)
	if reflectValue.Kind() == reflect.String {
		fmt.Println("value variabel: ", reflectValue.String()) // mengambil value strings
	}

	// jika hanya memerlukan nilai dari variabel, maka bisa menggunaka interface

	number = 90
	reflectValue = reflect.ValueOf(number)
	nilai := reflectValue.Interface()
	fmt.Println("nilai: ", nilai)

	// atau dapat langsung mengcasting interface nya
	nilai = reflectValue.Interface().(int)
	fmt.Println("nilai: ", nilai)
}

type Student struct {
	Name  string
	Grade int
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func IdentyfingMethod() {
	s1 := &Student{Name: "Andi", Grade: 2}

	fmt.Println("s1 name:", s1.Name)

	// value dari struct nya
	reflectValue := reflect.ValueOf(s1)
	method := reflectValue.MethodByName("SetName") // mengambil method sesuai namanya

	// menjalankan method
	method.Call([]reflect.Value{
		reflect.ValueOf("andi update"),
	})

	fmt.Println("1 name setelah update:", s1.Name)

}
