package main

import (
	"fmt"
)

type User struct {
	name    string
	age     int
	country string
}

func filter_words(input []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range input {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func filter_users(input []User, f func(User) bool) []User {
	filtered := make([]User, 0)
	for _, v := range input {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	// words := []string{"war", "water", "cup", "tree", "storm"}
	// p := "w"
	// result := filter_words(words, func(s string) bool {
	// 	return strings.HasPrefix(s, p)
	// })
	// fmt.Println(result)

	users := []User{
		{"John Doe", 40, "USA"},
		{"Hung Le", 22, "USA"},
		{"Paul Smith", 30, "Canada"},
		{"Lucia Mala", 34, "Slovakia"},
		{"Patrick Connor", 33, "USA"},
		{"Tim Welson", 55, "Canada"},
		{"Tomas Smutny", 38, "Slovakia"},
	}
	result := filter_users(users, func(u User) bool {
		return u.age > 30
	})
	fmt.Println(result)
}
