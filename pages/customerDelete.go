package pages

import (
	"db"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func CustomerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		customerID := r.FormValue("CustomerID")
		if customerID == "" {
			customerID = r.URL.Query().Get("CustomerID")
		}
		id, err := strconv.Atoi(customerID)
		if err != nil {
			// handle the error, e.g. return an error response
			return
		}
		log.Println("id:", id, "type:", reflect.TypeOf(id))
		db.DeleteCustomer(id)
		// if err != nil {
		// 	// handle the error, e.g. return an error response
		// 	return
		// }
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Customer deleted successfully!")
	}
}
