package services

import (
	"conectar_db_api/models"
	"database/sql"
	"fmt"
)

// GetItems obtiene todos los items de la tabla 'items' de la base de datos.
// Retorna una lista de struct 'models.Item' y un error en caso de que haya ocurrido alguno.
func GetItems() ([]models.Item, error) {
	items := make([]models.Item, 0)
	rows, err := Db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	// Itera sobre cada fila en 'rows' y crea una instancia de 'models.Item' con los valores de cada columna.
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

func GetItemById(id int) (*models.Item, error) {
	query := `
    SELECT id, customer_name, order_date, product, quantity, price
        FROM items WHERE id = $1;
    `
	var item models.Item

	err := Db.QueryRow(query, id).Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró ningún objeto
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func SearchItemByName(customerName string) ([]models.Item, error) {

	query := `
    SELECT id, customer_name, order_date, product, quantity, price
        FROM items WHERE LOWER(customer_name) = $1;
    `
	rows, err := Db.Query(query, customerName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return items, nil
}

func UpdateItemByID(id int, item models.Item) (models.Item, error) {
	var UpdateItem models.Item

	row := Db.QueryRow("UPDATE items SET CustomerName = $1, OrderDate = $2, Product = $3, Quantity = $4, Price = $5 WHERE id = $6 RETURNING *",
		item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, id)
	err := row.Scan(&UpdateItem.ID, &UpdateItem.CustomerName, &UpdateItem.OrderDate, &UpdateItem.Product, &UpdateItem.Quantity, &UpdateItem.Price)
	if err != nil {
		return models.Item{}, err
	}
	return UpdateItem, nil
}

func DeletItemByID(id int, item models.Item) (models.Item, error) {

	_, err := Db.Exec("DELETE FROM items WHERE id = $1 ",
		item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, item.Views, id)

	if err != nil {
		return models.Item{}, err
	}

	return item, nil

}

func CreateItem(item models.Item) (models.Item, error) {
	//item.CalculateTotalPrice()
	_, err := Db.Exec("INSERT INTO items (customer_name, order_date, product, quantity, price,views_counter)",
		item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, item.Views)

	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func ViewCounter(id int, mu *sync.Mutex) error {
	mu.Lock()

	query := "UPDATE items SET views_counter = views_counter + 1 WHERE id = $1"
	_, err := Db.Exec(query, id)
	if err != nil {
		return err
	}

	mu.Unlock()
	return nil
}

