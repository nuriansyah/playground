package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	//TODO: You must implement this function fot fetch product by id
	sqlStmt := `SELECT * FROM products WHERE id = ?`
	var product Product

	row := p.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&product.ID,
		&product.Price,
		&product.ProductName,
		&product.Quantity,
		&product.Category,
	)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	// TODO: You must implement this function for fetch product by name
	sqlStmt := `SELECT * FROM products WHERE product_name = ?`

	rows, err := p.db.Query(sqlStmt, productName)
	if err != nil {
		return Product{}, err
	}
	defer rows.Close()
	var product Product
	for rows.Next() {

		err := rows.Scan(
			&product.ID,
			&product.Price,
			&product.ProductName,
			&product.Quantity,
			&product.Category,
		)
		if err != nil {
			return Product{}, err
		}

	}

	return product, nil // TODO: replace this

}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	// TODO: You must implement this function for fetch all products
	sqlStmt := `SELECT * FROM products`
	rows, err := p.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	products := []Product{}
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Price,
			&product.ProductName,
			&product.Quantity,
			&product.Category,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil // TODO: replace this
}
