package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	Customer_ID  string `json:"customer_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_Number string `json:"phone_number"`
}

type Menu struct {
	Item_ID     string `json:"item_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type Orders struct {
	Order_ID   string    `json:"order_id"`
	CustomerID string    `json:"customer_id"`
	OrderTime  time.Time `json:"order_time"`
	PriceTotal int       `json:"price_total"`
}

type OrderItem struct {
	Order_ID  string `json:"order_id"`
	Item_ID   string `json:"item_id"`
	Quantity  int    `json:"quantity"`
	PriceUnit int    `json:"price_unit"`
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Connection() sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gorest")
	CheckError(err)
	defer db.Close()
	return *db
}
