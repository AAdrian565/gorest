package db

import (
	"fmt"
	"sort"
	"time"
)

type OrderInput struct {
	Item_ID  int
	Quantity int
}
type OrderDetails struct {
	OrderID      int
	Date         string
	CustomerName string
	Items        []Item
}

type Item struct {
	ItemName string
	Quantity int
	Price    float64
}

func SelectOrder() []OrderDetails {
	db := Connection()
	defer db.Close()
	query := `
SELECT o.Order_ID, o.Order_Time, c.Name, m.Name, oi.Quantity, m.Price
FROM Orders o
JOIN Customer c ON o.Customer_ID = c.Customer_ID
JOIN Order_Item oi ON o.Order_ID = oi.Order_ID
JOIN Menu m ON oi.Item_ID = m.Item_ID
ORDER BY o.Order_ID`

	rows, err := db.Query(query)
	CheckError(err)
	defer rows.Close()

	orderDetailsMap := make(map[int]OrderDetails)

	for rows.Next() {
		var orderID int
		var date, customerName, itemName string
		var quantity int
		var price float64

		err := rows.Scan(&orderID, &date, &customerName, &itemName, &quantity, &price)
		CheckError(err)

		item := Item{ItemName: itemName, Quantity: quantity, Price: price}

		if orderDetails, ok := orderDetailsMap[orderID]; ok {
			orderDetails.Items = append(orderDetails.Items, item)
			orderDetailsMap[orderID] = orderDetails
		} else {
			orderDetailsMap[orderID] = OrderDetails{OrderID: orderID, Date: date, CustomerName: customerName, Items: []Item{item}}
		}
	}

	var orderDetailsList []OrderDetails
	for _, orderDetails := range orderDetailsMap {
		orderDetailsList = append(orderDetailsList, orderDetails)
	}

	sort.Slice(orderDetailsList, func(i, j int) bool {
		return orderDetailsList[i].OrderID < orderDetailsList[j].OrderID
	})
	return orderDetailsList
}

func AddOrder(customerID int, orderItems []OrderInput) error {
	db := Connection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	orderTime := time.Now()
	res, err := tx.Exec("INSERT INTO Orders (Customer_ID, Order_Time) VALUES (?, ?)", customerID, orderTime)
	if err != nil {
		tx.Rollback()
		return err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range orderItems {
		var priceUnit int
		err := tx.QueryRow("SELECT Price FROM Menu WHERE Item_ID = ?", item.Item_ID).Scan(&priceUnit)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec("INSERT INTO Order_Item (Order_ID, Item_ID, Quantity, Price_Unit) VALUES (?, ?, ?, ?)", orderID, item.Item_ID, item.Quantity, priceUnit)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("Order %d added successfully\n", orderID)
	return nil
}
