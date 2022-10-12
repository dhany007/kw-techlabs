package models

type Employee struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Departement string `json:"departement"`
}

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}
