package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	ID					int			`json:"id"`
	ProductID		int			`json:"product_id"`
	CustomerID	int			`json:"customer_id"`
	Status			string	`json:"status"`
}

var orderID = 1

var orders = []Order{}

func main() {
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var order Order
			json.NewDecoder(r.Body).Decode(&order)
			order.ID = orderID
			orderID++
			orders = append(orders, order)
			json.NewEncoder(w).Encode(order)
		} else {
			json.NewEncoder(w).Encode(orders)
		}
	})

	fmt.Println("Starting server on port 8001")
	http.ListenAndServe(":8001", nil)
}