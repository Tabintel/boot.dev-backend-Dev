package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payment struct{
	ID				int			`json:"id"`
	OrderID		int			`json:"order_id"`
	Amount		int			`json:"amount"`
	Status		string	`json:"status"`
}

var paymentID = 1

var payments = []Payment{}

func main() {
	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var payment Payment
			json.NewDecoder(r.Body).Decode(&payment)
			payment.ID = paymentID
			paymentID++
			payment.Status = "success"
			payments = append(payments, payment)
			json.NewEncoder(w).Encode(payment)
		} else {
			json.NewEncoder(w).Encode(payments)
		}
	})

	// displays the server starting

	fmt.Println("Starting server on port 8002")
	http.ListenAndServe(":8002", nil)
}