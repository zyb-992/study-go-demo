package main

import "fmt"

// deferWithClosure should print 2
func deferWithClosure() {
	result := 1
	defer func() {
		fmt.Println("deferWithClosure", result)
	}()
	result = 2
}

// deferWithClosurePassParam should print 1
func deferWithClosurePassParam() {
	result := 1
	defer func(i int) {
		fmt.Println("deferWithClosure", i)
	}(result)
	result = 2
}

// deferWithoutClosure should print 1
func deferWithoutClosure() {
	result := 1
	defer fmt.Println("deferWithoutClosure", result)
	result = 2
}

// deferWithClosureUsingConcreteName should print 2
func deferWithClosureUsingConcreteName() (result int) {
	result = 1
	defer func() {
		result = 2
	}()

	return result
}

// deferWithClosureUsingAnonymousName should print 1
func deferWithClosureUsingAnonymousName() int {
	result := 1
	defer func() {
		result = 2
	}()

	return result
}
