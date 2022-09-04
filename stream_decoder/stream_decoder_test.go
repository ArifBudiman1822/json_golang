package streamdecoder

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Address struct {
	City     string `json:"city"`
	Province string `json:"province"`
}

type Customer struct {
	Name    string    `json:"name"`
	Address []Address `json:"address`
}

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("customer.json")

	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}
