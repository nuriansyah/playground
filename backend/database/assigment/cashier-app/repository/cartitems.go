package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	//TODO: add sql statement here
	//HINT: join table cart_items and products

	sqlStatement = `SELECT * FROM cart_items JOIN products ON cart_items.product_id = products.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(
			&cartItem.ProductID,
			&cartItem.ProductName,
			&cartItem.Quantity,
			&cartItem.Price,
			&cartItem.Category,
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement

	sqlStatement = `SELECT * FROM cart_items WHERE product_id = ?`

	row := c.db.QueryRow(sqlStatement, productID)
	err := row.Scan(&cartItem.ProductID, &cartItem.ProductName, &cartItem.Quantity, &cartItem.Price, &cartItem.Category)
	if err != nil {
		return CartItem{}, err
	}
	return cartItem, nil
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item
	sqlStmt := `INSERT INTO
	cart_items (product_id, quantity, price, ProductName, Category)
	VALUES 
	(?, ?, ?,?, ?)`
	_, err := c.db.Exec(sqlStmt, cartItem.Quantity, cartItem.Price, cartItem.Category, cartItem.ProductName)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	//HINT : you can use the update statement

	sqlStmt := `UPDATE cart_items SET quantity + ? WHERE product_id = ?`

	_, err := c.db.Exec(sqlStmt, cartItem.Quantity, cartItem.ProductID)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement

	sqlStmt := `DELETE FROM cart_items`

	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement

	sqlStatement = `SELECT SUM(p.price * c.quantity ) FROM cart_items c LEFT JOIN products p ON c.product_id = products.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var totalPrice int
	for rows.Next() {
		var price int
		err := rows.Scan(&price)
		if err != nil {
			return 0, err
		}
		totalPrice += price
	}
	return totalPrice, nil
}
