package main

import (
	"fmt"
	"strings"
)

func common() {
	s := "learning go \u2318"
	sub := "learn"

	// should be `learning go ⌘`
	cloneS := strings.Clone(s)
	fmt.Printf("strings.Clone -> clone str:%s, origin str:%s\n", cloneS, s)

	// should be true
	bool1 := strings.ContainsAny(s, sub)
	fmt.Printf("strings.ContainsAny -> bool:%v\n", bool1)

	// should be true
	bool2 := strings.ContainsRune(s, '⌘')
	fmt.Printf("strings.ContainsRune -> bool:%v\n", bool2)

	index := strings.Index(s, sub)
	fmt.Printf("strings.Index -> index:%v\n", index)

	// should be 2
	cnt := strings.Count(s, "n")
	fmt.Printf("strings.Count ->  count:%v\n", cnt)

}
