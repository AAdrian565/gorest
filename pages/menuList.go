package pages

import (
	"db"
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
)

func ListMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	menuItems := db.SelectMenu()
	acceptHeader := r.Header.Get("Accept")

	// web
	if strings.Contains(acceptHeader, "text/html") {
		tmpl := template.Must(template.ParseFiles("pages/view/menuList.html"))
		err := tmpl.Execute(w, menuItems)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	jsonData, err := json.Marshal(menuItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
