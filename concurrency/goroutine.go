package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func print_m() {
	numbers()
	alphabets()
}

func print_go() { // print with go routine
	go numbers()
	go alphabets()
	time.Sleep(100 * time.Millisecond) // withour sleep, main function terminated before goroutine print
}
func main() {
	print_m()
	// print_go()
	fmt.Println("main terminated")
}
