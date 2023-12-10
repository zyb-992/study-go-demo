package main

import "reflect"

func main() {

}

func IfEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
