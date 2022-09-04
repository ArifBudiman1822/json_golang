package test_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {

	customer := Customer{
		Firstname:  "Muhammad",
		Middlename: "Arif",
		Lastname:   "Budiman",
		Age:        21,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}
