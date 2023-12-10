package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string `json:"name"`
}

func main() {
	jsonData := `{"name":"zhengyb","age":123}`
	d := Data{}
	err := json.Unmarshal([]byte(jsonData), &d)
	fmt.Println(err)
	fmt.Println(d)
}
