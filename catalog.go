package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID		int			`json:"id"`
	Name	string	`json:"name"`
	Price	int			`json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Product 1", Price: 100},
	{ID: 2, Name: "Product 2", Price: 200},
	{ID: 3, Name:	"Product 3", Price: 300},
}

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(products)
	})

	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", nil)
}