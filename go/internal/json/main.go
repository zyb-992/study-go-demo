package main

import (
	"encoding/json"
	"fmt"
)

type jsonData struct {
	Name string `json:"name"`
	Str  string `json:"str"`
}

func main() {
	const str = `{"name":"zib","age":"twenty"}`

	a := jsonData{
		Name: "zhengyb",
		Str:  str,
	}

	data, _ := json.Marshal(a)
	fmt.Println(string(data))
	var v jsonData
	json.Unmarshal(data, &v)
	fmt.Printf("%+v", v)
}
