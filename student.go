package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

// Person can use as embed struct
// Person and int are anonymous field
type Student struct {
	Person // user-defined type
	int    // built-in type
	addr   string
}

type Student2 struct {
	*Person
	id   int
	addr string
}

func main() {
	s1 := Student{Person{"5lmh", "man", 20}, 1, "bj"}
	fmt.Println(s1)

	s2 := Student{Person: Person{"5lmh", "man", 20}}
	fmt.Println(s2)

	st_pointer := Student2{&Person{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(st_pointer)
}
