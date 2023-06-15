package main

import (
	"db"
	"encoding/json"
	"net/http"
	"pages"
	"web"
)

func displayData(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listCustomers(w http.ResponseWriter, r *http.Request) {
	customers := db.SelectCustomer()
	displayData(w, customers)
}

func main() {

	// orderDetailsList := db.SelectOrder()
	// for _, orderDetails := range orderDetailsList {
	// 	fmt.Printf("Order ID: %d, Date: %s, Customer Name: %s\n", orderDetails.OrderID, orderDetails.Date, orderDetails.CustomerName)
	// 	for _, item := range orderDetails.Items {
	// 		fmt.Printf("  Item Name: %s, Quantity: %d, Price: %.2f\n", item.ItemName, item.Quantity, item.Price)
	// 	}
	// }
	http.HandleFunc("/order/input", pages.OrderInput)
	http.HandleFunc("/order", pages.OrderList)

	http.HandleFunc("/menu/input", pages.MenuInput)
	http.HandleFunc("/menu/edit", pages.MenuEdit)
	http.HandleFunc("/menu/delete", pages.MenuDelete)
	http.HandleFunc("/menu", pages.ListMenu)

	http.HandleFunc("/customer/input", pages.CustomerInput)
	http.HandleFunc("/customer/edit", pages.CustomerEdit)
	http.HandleFunc("/customer/delete", pages.CustomerDelete)
	http.HandleFunc("/customer", pages.CustomerList)

	http.HandleFunc("/", web.ServeIndex)
	http.ListenAndServe(":8080", nil)
}

// db.AddCustomer("test123", "email@mail.com", "123123123")
