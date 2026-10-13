package main

import (
	"fmt"
	"time"
)

func decorator(f func() string) func() {
	return func() {
		start := time.Now()
		result := f()
		fmt.Printf("Execution took: %v\n", time.Since(start))
		fmt.Println(result)
	}
}

func myFunction() string {
	time.Sleep(100 * time.Millisecond)
	return "Function executed"
	
}

func main() {
	decoratedFunction := decorator(myFunction)
	decoratedFunction()
}
