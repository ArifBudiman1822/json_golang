package jsontags

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Products struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTags(t *testing.T) {

	products := Products{
		Id:       "P001",
		Name:     "Makanan",
		ImageURL: "https://domain.com/image_url",
	}

	bytes, _ := json.Marshal(products)

	fmt.Println(string(bytes))

}

func TestJsonTagsDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Makanan","image_url":"https://domain.com/image_url"}`

	jsonBytes := []byte(jsonString)

	product := &Products{}

	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
