package pages

import (
	"db"
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
	"web"
)

func CustomerList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	customers := db.SelectCustomer()
	acceptHeader := r.Header.Get("Accept")

	// web
	if strings.Contains(acceptHeader, "text/html") {
		tmpl := template.Must(template.ParseFiles("pages/view/customerList.html"))
		err := tmpl.Execute(w, customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	jsonData, err := json.Marshal(customers)
	web.ErrorCheck(w, r, err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
