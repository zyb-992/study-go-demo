package main

import "fmt"

func main() {
	usingClosure()

	deferWithClosure()
	deferWithoutClosure()
	deferWithClosurePassParam()

	res := deferWithClosureUsingConcreteName()
	fmt.Println("deferWithClosureUsingConcreteName", res)
	res = deferWithClosureUsingAnonymousName()
	fmt.Println("deferWithClosureUsingAnonymousName", res)
}
