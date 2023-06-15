package pages

import (
	"db"
	"fmt"
	"net/http"
	"web"
)

func CustomerInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Name := r.FormValue("Name")
		Email := r.FormValue("Email")
		Phone_Number := r.FormValue("Phone_Number")

		if Name == "" && Email == "" && Phone_Number == "" {
			Name = r.URL.Query().Get("Name")
			Email = r.URL.Query().Get("Email")
			Phone_Number = r.URL.Query().Get("Phone_Number")
		}

		// Insert the customer into the database
		db.AddCustomer(Name, Email, Phone_Number)

		w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, "Customer added successfully!")
		return
	}
	web.Serve("customerInput")
	fmt.Fprintln(w, `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Form Example</title>
	<style>
	body {
		font-family: Arial, sans-serif;
		margin: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	.container {
		max-width: 400px;
		width: 100%;
		padding: 20px;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
	}

	h1 {
		text-align: center;
	}

	label {
		display: block;
		margin-bottom: 5px;
	}

	input[type="text"],
	input[type="email"],
	input[type="tel"] {
		width: 100%;
		padding: 8px;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
		margin-bottom: 10px;
	}

	.submit-button {
		display: flex;
		justify-content: flex-end;
	}

	.submit-button input[type="submit"] {
		background-color: #4CAF50;
		color: white;
		padding: 10px 20px;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		font-size: 16px;
	}

	.submit-button input[type="submit"]:hover {
		background-color: #45a049;
	}
	</style>
	</head>
	<body>
	<div class="container">
	<h1>Add new customer:</h1>
	<form method="POST">
	<label for="name">Name:</label>
	<input type="text" id="Name" name="Name" required>

	<label for="email">Email:</label>
	<input type="email" id="Email" name="Email" required>

	<label for="phone">Phone Number:</label>
	<input type="tel" id="Phone_Number" name="Phone_Number" required>

	<div class="submit-button">
	<input type="submit" value="Submit">
	</div>
	</form>
	</div>
	</body>
	</html>
	`)
}
