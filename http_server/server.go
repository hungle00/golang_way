package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Product struct {
	Name         string
	Price        string
	Description  string
	ShippingDate time.Time
}

// Rendering a HTML Template
func productHandler(w http.ResponseWriter, r *http.Request) {

	html, err := template.ParseFiles("views/product.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	product := Product{Name: "Red Tea Pot 250ml", Description: "Test", Price: "19.99", ShippingDate: time.Now()}
	err = html.Execute(w, product)
}

// Rendering JSON
func productJsonHandler(w http.ResponseWriter, r *http.Request) {
	product := Product{Name: "Red Tea Pot 250ml", Description: "Test", Price: "19.99", ShippingDate: time.Now()}
	json, err := json.Marshal(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product Handler")
		message := []byte("Welcome to product!")
		w.Write(message)
	})

	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/product.json", productJsonHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
