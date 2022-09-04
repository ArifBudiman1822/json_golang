package test_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logjson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestEncodeJson(t *testing.T) {

	logjson("Arif")
	logjson(26)
	logjson(true)
	logjson(false)
	logjson([]string{"Muhammad", "Arif", "Budiman"})
}
