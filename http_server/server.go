package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"log"
	"os"
)

type Product struct {
	Name         string
	Price        string
	Description  string
	ShippingDate time.Time
}

var logger = log.New(os.Stdout, "", 0)

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

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // ghi nhận thời gian trước khi chạy
        timeStart := time.Now()

        // next là hàm business logic được truyền vào
        next.ServeHTTP(w, r)
        // tính toán thời gian thực thi
        timeElapsed := time.Since(timeStart)
        // log ra thời gian thực thi
        logger.Println(timeElapsed)
    })
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product Handler")
		message := []byte("Welcome to product!")
		w.Write(message)
	})
	http.Handle("/hello", timeMiddleware(http.HandlerFunc(helloHandler)))
	http.Handle("/product.json", timeMiddleware(http.HandlerFunc(productJsonHandler)))

	http.HandleFunc("/product", productHandler)
	// http.HandleFunc("/product.json", productJsonHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
