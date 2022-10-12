package struct_go

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Division string
	Address  Address
}

type Address struct {
	City  string
	Jalan string
}

type Person struct {
	Name string
	Age  int
}

func GivingValueStruct() {
	employee := Employee{}
	employee.Name = "Dhany"
	employee.Age = 20
	employee.Division = "IT"

	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Division:", employee.Division)

	fmt.Println("\ninitializing struct")
	employee2 := Employee{}
	employee2.Name = "Kalai"
	employee2.Age = 20
	employee2.Division = "BD"
}

func InitializingStruct() {
	employee2 := Employee{}
	employee2.Name = "Kalai"
	employee2.Age = 20
	employee2.Division = "BD"

	employee3 := Employee{Name: "jhon", Age: 20, Division: "Sales"}
	fmt.Printf("Employee1 : %+v\n", employee2)
	fmt.Printf("Employee2 : %+v\n", employee3)

}

func PointerToStruct() {
	employee3 := Employee{Name: "jhon", Age: 20, Division: "Sales"}

	employee4 := &employee3

	fmt.Println("employe3 name: ", employee3.Name)
	fmt.Println("employe4 name: ", employee4.Name)

	employee4.Name = "dhany"
	fmt.Println("after employe3 name: ", employee3.Name)
	fmt.Println("after employe4 name: ", employee4.Name)
}

func EmbeddedStruct() {
	employeeAni := Employee{}
	employeeAni.Name = "Ani"
	employeeAni.Age = 20
	employeeAni.Division = "Teacher"
	employeeAni.Address.City = "Jambi"
	employeeAni.Address.Jalan = "Gudang Arang"
	fmt.Printf("Employe : %+v\n", employeeAni)
}

func AnynomousStruct() {
	// anynomous struct : langsung dideklarasikan bersamaan dengan pembuatan variabel nya
	// tanpa pengisian field
	employeeKalai := struct {
		Person   Person
		Division string
	}{}
	employeeKalai.Person.Name = "Kalai"
	employeeKalai.Person.Age = 20
	employeeKalai.Division = "Teacher"
	fmt.Printf("employe : %+v\n", employeeKalai)

	// dengan pengisian field
	employeeHani := struct {
		Person   Person
		Division string
	}{
		Person:   Person{Name: "Hani", Age: 20},
		Division: "Marketing",
	}
	fmt.Printf("employe : %+v\n", employeeHani)

}

func SliceOfStruct() {
	persons := []Person{
		{Name: "arif", Age: 20},
		{Name: "Budi", Age: 20},
		{Name: "ani", Age: 20},
	}

	for _, person := range persons {
		fmt.Printf("%+v\n", person)
	}
}

func SliceAnonymousStruct() {
	employe := []struct {
		Person   Person
		Division string
	}{
		{Person: Person{Name: "Ani", Age: 20}, Division: "IT"},
		{Person: Person{Name: "Ane", Age: 20}, Division: "IT"},
		{Person: Person{Name: "Ana", Age: 20}, Division: "IT"},
	}

	for _, v := range employe {
		fmt.Printf("%+v\n", v)
	}
}
