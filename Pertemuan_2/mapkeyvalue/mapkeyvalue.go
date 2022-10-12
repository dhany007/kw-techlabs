package mapkeyvalue

import "fmt"

func MapKeyValue() {
	fmt.Println("map")

	var personInisialisasi = map[string]string{}

	personInisialisasi["name"] = "dhany"
	personInisialisasi["age"] = "20"
	personInisialisasi["address"] = "jakarta"

	fmt.Println(personInisialisasi["name"])
	fmt.Println(personInisialisasi["age"])
	fmt.Println(personInisialisasi["address"])

	var person = map[string]string{
		"name":    "kalai",
		"age":     "18",
		"address": "simanabun",
	}

	fmt.Println("before deleting: ", person)

	delete(person, "address")

	fmt.Println("after deleting address: ", person)

	for k, v := range person {
		fmt.Println("key : ", k, " value: ", v)
	}

	var persons = []map[string]string{
		{
			"name": "dhany",
			"age":  "20",
		},
		{
			"name": "kalai",
			"age":  "30",
		},
	}

	for i, person := range persons {
		fmt.Println("index: ", i, "name: ", person["name"], " age: ", person["age"])
	}

	fmt.Println("detecting value ")

	value, exist := person["fullname"]

	if exist {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("value telah didelete")
	}
}
