package pages

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderInputJson struct {
	CustomerID int             `json:"customer_id"`
	Items      []db.OrderInput `json:"items"`
}

func OrderInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var input OrderInputJson
		err := decoder.Decode(&input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = db.AddOrder(input.CustomerID, input.Items)
		if err != nil {
			http.Error(w, "Error adding order: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Order added successfully!")
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
