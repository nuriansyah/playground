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

	sqlStatement = `
	SELECT
		cart_items.id,
		cart_items.product_id,
		cart_items.quantity,
		products.product_name,
		products.price
	FROM cart_items
	JOIN products
	ON cart_items.product_id = products.id
	`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.ProductID,
			&cartItem.Quantity,
			&cartItem.ProductName,
			&cartItem.Price,
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil

	// return []CartItem{}, nil // TODO: replace this
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement
	sqlStatement = `
	SELECT
		id, 
		quantity
	FROM cart_items
	WHERE product_id = ?`

	row := c.db.QueryRow(sqlStatement, productID)
	if err := row.Scan(&cartItem.ID, &cartItem.Quantity); err != nil {
		return cartItem, err
	}

	return cartItem, nil

	// return CartItem{}, nil // TODO: replace this
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item
	sqlStatement := `
	INSERT INTO cart_items (
		product_id,
		quantity
	) VALUES (?, ?)`

	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}

	return nil
	// return nil // TODO: replace this
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	sqlStatement := `
	UPDATE cart_items
	SET 
	quantity = quantity+?
	WHERE id = ?`

	_, err := c.db.Exec(sqlStatement, cartItem.Quantity, cartItem.ID)
	if err != nil {
		return err
	}

	return nil

	// return nil // TODO: replace this
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement
	sqlStatement := `
	DELETE FROM cart_items WHERE id = ?`

	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil

	// return nil // TODO: replace this
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement
	sqlStatement = `
	SELECT SUM(p.price * c.quantity) 
	FROM cart_items c
	LEFT JOIN products p
	ON c.product_id = p.id
	`
	row, err := c.db.Query(sqlStatement)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	var totalPrice int
	for row.Next() {
		err := row.Scan(&totalPrice)
		if err != nil {
			return 0, err
		}
	}

	return totalPrice, nil
	// return 0, nil // TODO: replace this
}
