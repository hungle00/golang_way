// Closures are a special case of anonymous functions. 
// Closures are anonymous functions that access the variables defined outside the body of the function.
// Example from https://go.dev/tour/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		a := x
		x, y = y, x + y
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
