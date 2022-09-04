package jsonarray

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	City       string
	Postalcode int
}

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Married    bool
	Hobbies    []string
	Address    []Address
}

func TestJsonArray(t *testing.T) {

	customer := Customer{
		Firstname:  "Muhammad",
		Middlename: "Arif",
		Lastname:   "Budiman",
		Age:        21,
		Married:    false,
		Hobbies:    []string{"Swimming", "Running", "Coding"},
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {

	jsonString := `{"Firstname":"Muhammad","Middlename":"Arif","Lastname":"Budiman","Age":21,"Married":false,"Hobbies":["Swimming","Running","Coding"]}`

	jsonByte := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonByte, &customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)

}

func TestJsonArrayComplex(t *testing.T) {

	customer := Customer{
		Firstname: "Arif",
		Address: []Address{
			{
				Street:     "Ngampelsari",
				City:       "Sidoarjo",
				Postalcode: 61271,
			},
			{
				Street:     "Sumorame",
				City:       "Sidoarjo",
				Postalcode: 61271,
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))

}

func TestJsonArrayComplexDecode(t *testing.T) {

	jsonString := `{"Firstname":"Arif","Middlename":"","Lastname":"","Age":0,"Married":false,"Hobbies":null,"Address":[{"Street":"Ngampelsari","City":"Sidoarjo","Postalcode":61271},{"Street":"Sumorame","City":"Sidoarjo","Postalcode":61271}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Address)
}

func TestArrayDecode(t *testing.T) {
	jsonArray := `[{"Street":"Ngampelsari","City":"Sidoarjo","Postalcode":61271},{"Street":"Sumorame","City":"Sidoarjo","Postalcode":61271}]`

	jsonBytes := []byte(jsonArray)

	address := &[]Address{}

	err := json.Unmarshal(jsonBytes, address)

	if err != nil {
		panic(err)
	}

	fmt.Println(address)
}

func TestArrayEncode(t *testing.T) {

	address := []Address{
		{
			Street:     "Ngampelsari",
			City:       "Sidoarjo",
			Postalcode: 61271,
		},
		{
			Street:     "Sumorame",
			City:       "Sidoarjo",
			Postalcode: 61271,
		},
	}

	bytes, _ := json.Marshal(address)

	fmt.Println(string(bytes))

}
