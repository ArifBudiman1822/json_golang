package json_object_ecode

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
	Hobbies    []string
}

func TestDecodeJSON(t *testing.T) {

	jsonString := `{"Firstname":"Muhammad","Middlename":"Arif","Lastname":"Budiman","Age":21,"Married":false}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)

}
