package pages

//
// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"
//
// 	_ "github.com/go-gota/gota/dataframe"
// 	_ "github.com/go-gota/gota/series"
// 	_ "github.com/go-sql-driver/mysql"
// )
//
// func main() {
// 	recommendedItem, err := recommendItem()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if recommendedItem != "" {
// 		fmt.Printf("The recommended item for this hour is: %s\n", recommendedItem)
// 	} else {
// 		fmt.Println("No recommendation available for this hour.")
// 	}
// }
//
// func recommendItem() (string, error) {
// 	// Connect to the database
// 	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/gorest")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer db.Close()
//
// 	query := `
// SELECT m.Name, o.Order_Time
// FROM Orders o
// JOIN Order_Item oi ON o.Order_ID = oi.Order_ID
// JOIN Menu m ON oi.Item_ID = m.Item_ID`
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer rows.Close()
//
// 	var itemNames []string
// 	var orderHours []float64
// 	for rows.Next() {
// 		var itemName string
// 		var orderTime time.Time
// 		err := rows.Scan(&itemName, &orderTime)
// 		if err != nil {
// 			return "", err
// 		}
// 		itemNames = append(itemNames, itemName)
// 		orderHours = append(orderHours, float64(orderTime.Hour()))
// 	}
//
// 	df := dataframe.New(
// 		series.New(orderHours, series.Float, "OrderHour"),
// 	)
//
// 	k := 24
// 	km := gotaclustering.NewKMeans(k, 100, df)
// 	km.Fit()
//
// 	currentHour := float64(time.Now().Hour())
// 	cluster := km.Predict(dataframe.New(series.New([]float64{currentHour}, series.Float, "OrderHour")))
// 	itemCounts := make(map[string]int)
// 	for i, c := range km.Labels {
// 		if c == cluster[0] {
// 			itemCounts[itemNames[i]]++
// 		}
// 	}
//
// 	// Return the most popular item as the recommendation
// 	var recommendedItem string
// 	maxCount := 0
// 	for item, count := range itemCounts {
// 		if count > maxCount {
// 			maxCount = count
// 			recommendedItem = item
// 		}
// 	}
// 	return recommendedItem, nil
// }
