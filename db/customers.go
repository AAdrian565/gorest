package db

import "fmt"

func SelectCustomer() []Customer {
	db := Connection()
	result, err := db.Query("SELECT * FROM Customer")
	CheckError(err)
	var userArray []Customer
	var usr = Customer{}
	for result.Next() {
		err := result.Scan(&usr.Customer_ID, &usr.Name, &usr.Email, &usr.Phone_Number)
		CheckError(err)
		userArray = append(userArray, Customer{Customer_ID: usr.Customer_ID, Name: usr.Name, Email: usr.Email, Phone_Number: usr.Phone_Number})
	}
	return userArray
}

func AddCustomer(name, email, phoneNumber string) error {
	db := Connection()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Customer (Name, Email, Phone_Number) VALUES (?, ?, ?)")
	CheckError(err)
	defer insert.Close()

	_, err = insert.Exec(name, email, phoneNumber)
	CheckError(err)

	return nil
}

func DeleteCustomer(Customer_ID int) error {
	db := Connection()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM Customer WHERE Customer_ID = ?")

	fmt.Printf("test\n")
	// fmt.Fprintf(w, "wrong access")
	CheckError(err)
	defer delete.Close()

	fmt.Printf("test2\n")
	_, err = delete.Exec(Customer_ID)
	CheckError(err)

	return nil
}

func EditCustomer(customerID int, name, email, phoneNumber string) error {
	db := Connection()
	defer db.Close()

	update, err := db.Prepare("UPDATE Customer SET Name = ?, Email = ?, Phone_Number = ? WHERE Customer_ID = ?")
	CheckError(err)
	defer update.Close()

	_, err = update.Exec(name, email, phoneNumber, customerID)
	CheckError(err)

	return nil
}
