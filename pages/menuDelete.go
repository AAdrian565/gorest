package pages

import (
	"db"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func MenuDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		itemID := r.FormValue("ItemID")
		if itemID == "" {
			itemID = r.URL.Query().Get("ItemID")
		}
		id, err := strconv.Atoi(itemID)
		if err != nil {
			// handle the error, e.g. return an error response
			return
		}
		log.Println("id:", id, "type:", reflect.TypeOf(id))
		db.DeleteMenu(id)
		// if err != nil {
		// 	// handle the error, e.g. return an error response
		// 	return
		// }
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Customer deleted successfully!")
	}
}
