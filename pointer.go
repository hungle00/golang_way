package main

import (
	"fmt"
)

// pass by pointer
func change(val *int) {
	*val = 55
}

// pass by value
func change1(val int) {
	val = 10
}

// pass a pointer to an array
func modify(arr *[3]int) {
	arr[1] = 5
}

func main() {
	a := 58
	change1(a)
	fmt.Println("value of a before function call is", a)
	fmt.Println("address of a is", &a)
	change(&a)
	fmt.Println("value of a after function call is", a)

	arr := [3]int{1, 6, 8}
	modify(&arr)
	fmt.Println(arr)
}
