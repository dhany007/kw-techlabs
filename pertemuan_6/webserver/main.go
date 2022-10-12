package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"webserver/datas"
	"webserver/models"
	"webserver/params"
)

func main() {
	fmt.Println("webserver golang")

	PORT := ":8080"

	http.HandleFunc("/employees", getEmployee)
	http.HandleFunc("/employee/create", createEmploye)
	http.HandleFunc("/employee/update", updateEmployee)
	http.HandleFunc("/employee/delete", deleteEmployee)

	fmt.Println("server running in port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	query := r.URL.Query()
	nameSearch := query.Get("name")

	w.Header().Add("Content-Type", "application/json")
	if nameSearch == "" {
		response := models.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Payload: datas.Employees,
		}
		responseEncode(w, response)
		return
	}

	emp := models.Employee{}

	for _, v := range datas.Employees {
		if strings.EqualFold(nameSearch, v.Name) {
			emp = v
			break
		}
	}

	if emp.Name == "" {
		response := models.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		}
		responseEncode(w, response)
		return
	}

	response := models.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Payload: emp,
	}
	responseEncode(w, response)
}

func createEmploye(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Add("Content-Type", "application/json")

	if method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	request := params.CreateEmployee{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response := models.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}
		responseEncode(w, response)
		return
	}

	emp := models.Employee{
		Name:        request.Name,
		Age:         request.Age,
		Departement: request.Departement,
	}

	datas.Employees = append(datas.Employees, emp)

	response := models.Response{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Payload: emp,
	}

	responseEncode(w, response)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Add("Content-Type", "application/json")

	if method != http.MethodPut {
		methodNotAllowed(w)
		return
	}

	request := params.CreateEmployee{}

	query := r.URL.Query()
	nameSearch := query.Get("name")

	indexEmployee := -1 // out of index (not found in array)

	for i, v := range datas.Employees {
		if strings.EqualFold(nameSearch, v.Name) && v.Name != "" {
			indexEmployee = i
			break
		}
	}

	if indexEmployee == -1 {
		response := models.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		}
		responseEncode(w, response)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response := models.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}
		responseEncode(w, response)
		return
	}

	emp := models.Employee{
		Name:        request.Name,
		Age:         request.Age,
		Departement: request.Departement,
	}
	datas.Employees[indexEmployee] = emp

	response := models.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Payload: emp,
	}

	responseEncode(w, response)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Add("Content-Type", "application/json")

	if method != http.MethodDelete {
		methodNotAllowed(w)
		return
	}

	query := r.URL.Query()
	nameSearch := query.Get("name")

	indexEmployee := -1 // out of index (not found in array)

	for i, v := range datas.Employees {
		if strings.EqualFold(nameSearch, v.Name) && v.Name != "" {
			indexEmployee = i
			break
		}
	}

	if indexEmployee == -1 {
		response := models.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		}
		responseEncode(w, response)
		return
	}

	/*
		arr = [1, 2, 3, 4]
		delete 2
		tempArr = []
		tempArr = [1] + [3,4]
	*/
	tempEmployees := append(datas.Employees[:indexEmployee], datas.Employees[indexEmployee+1:]...)
	datas.Employees = tempEmployees

	response := models.Response{
		Code:   http.StatusOK,
		Status: "OK",
	}

	responseEncode(w, response)
}

func methodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Add("Content-Type", "application/json")

	response := models.Response{
		Code:   http.StatusMethodNotAllowed,
		Status: "METHOD NOT ALLOWED",
	}
	json.NewEncoder(w).Encode(response)
}

func responseEncode(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
