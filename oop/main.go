package main

import "oop/employee"

func main() {
	e := employee.New("Steve", "Lee", 30, 20)
	e.LeavesRemaining()
}
