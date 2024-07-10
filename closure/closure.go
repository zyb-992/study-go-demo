package main

import "fmt"

/*
  all the anonymous function in golang is called "closure"
*/

// closureWithVar
func closureWithVar(count int) func() {
	return func() {
		count++
		fmt.Println(count)
	}
}

func closureWithPointer(count *int) func() {
	return func() {
		*count++
		fmt.Println(*count)
	}
}

func usingClosure() {
	i := 0
	f := closureWithVar(i)
	// first call
	f()
	// second call
	f()
	// third call
	f()

	i = 10
	// what will print after exec f() ?
	f()

	pi := 0
	pf := closureWithPointer(&pi)
	// first call
	pf()
	// second call
	pf()

	// change val that pointer points to
	pi = 10

	// what will print after exec pf() ?
	pf()

}
