package pages

import (
	"db"
	"fmt"
	"net/http"
	"strconv"
)

func MenuEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		ID := r.FormValue("ID")
		Name := r.FormValue("Name")
		Description := r.FormValue("Description")
		PriceStr := r.FormValue("Price")

		if Name == "" && Description == "" && PriceStr == "" {
			ID = r.URL.Query().Get("ID")
			Name = r.URL.Query().Get("Name")
			Description = r.URL.Query().Get("Description")
			PriceStr = r.URL.Query().Get("Price")
		}

		Price, _ := strconv.Atoi(PriceStr)
		menuID, _ := strconv.Atoi(ID)
		db.EditMenu(menuID, Name, Description, Price)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Menu editet successfully")
		return
	}
	// web.Serve("customerInput")
	fmt.Fprintln(w, `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Form Example</title>
	</head>
	<body>
	<h1>Fill out the form:</h1>
	<form method="POST">
	<label for="name">Name:</label>
	<input type="text" id="Name" name="Name" required><br><br>

	<label for="email">Email:</label>
	<input type="email" id="Email" name="Email" required><br><br>

	<label for="phone">Phone Number:</label>
	<input type="tel" id="Phone_Number" name="Phone_Number" required><br><br>

	<input type="submit" value="Submit">
	</form>
	</body>
	</html>
	`)
}
