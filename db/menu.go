package db

func SelectMenu() []Menu {
	db := Connection()
	result, err := db.Query("SELECT * FROM Menu")
	CheckError(err)
	var menuArray []Menu
	var menu = Menu{}
	for result.Next() {
		err := result.Scan(&menu.Item_ID, &menu.Name, &menu.Description, &menu.Price)
		CheckError(err)
		menuArray = append(menuArray, Menu{Item_ID: menu.Item_ID, Name: menu.Name, Description: menu.Description, Price: menu.Price})
	}
	return menuArray
}

func AddMenuItem(name string, description string, price int) error {
	db := Connection()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Menu (Name, Description, Price) VALUES (?, ?, ?)")
	CheckError(err)
	defer insert.Close()

	_, err = insert.Exec(name, description, price)
	CheckError(err)

	return nil
}

func DeleteMenu(Item_ID int) error {
	db := Connection()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM Menu WHERE Item_ID = ?")
	CheckError(err)
	defer delete.Close()

	_, err = delete.Exec(Item_ID)
	CheckError(err)

	return nil
}
func EditMenu(Item_ID int, name string, description string, price int) error {
	db := Connection()
	defer db.Close()

	update, err := db.Prepare("UPDATE Menu SET Name = ?, description = ?, price = ? WHERE Item_ID = ?")
	CheckError(err)
	defer update.Close()

	_, err = update.Exec(name, description, price, Item_ID)
	CheckError(err)

	return nil
}
