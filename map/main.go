package main

import "fmt"

func main() {
	variable := map[string]interface{}{}
	buyList := make([]map[string]interface{}, 2)
	map1 := map[string]interface{}{}
	map2 := map[string]interface{}{}
	map1["wait_pay_order_id"] = 23
	map2["wait_pay_order_id"] = 24

	buyList[0] = map1
	buyList[1] = map2

	variable["buy_list"] = buyList

	_, ok := variable["buy_list"].([]interface{})
	if !ok {
		fmt.Println("type assert Error")
	}

	fmt.Printf("buy_list type: %T value: %+v", variable["buy_list"], variable["buy_list"])

}
