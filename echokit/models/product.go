package models

import (
	"errors"
	"fmt"
	"time"
)

type (
	// Product model
	Product struct {
		ProductID    string    `json:"product_id" sql:""`
		ProductName  string    `json:"product_name" sql:""`
		ProductDesc  string    `json:"product_desc" sql:""`
		ProductImage string    `json:"product_image" sql:""`
		ProductPrice float64   `json:"product_price" sql:""`
		CreatedAt    time.Time `json:"created_at" sql:"default:now()"`
		UpdatedAt    time.Time `json:"updated_at" sql:"default:now()"`
	}
)

// GetProduct to getting products
func (m *MDL) GetProduct() ([]Product, error) {
	db := m.MysqlClient

	qs := `
	SELECT
	* FROM product;`

	rows, err := db.Query(qs)
	if err != nil {
		return []Product{}, err
	}
	rowIndex := 0
	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.ProductDesc,
			&product.ProductImage,
			&product.ProductPrice,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return []Product{}, err
		}
		products = append(products, product)
		rowIndex++
	}

	return products, nil
}

// GetProductByID to getting product by id
func (m *MDL) GetProductByID(id string) (Product, error) {
	db := m.MysqlClient
	// Execute the query
	qs := "SELECT * FROM product WHERE product_id = ?;"

	rows, err := db.Query(qs, id)
	if err != nil {
		return Product{}, err
	}
	rowIndex := 0
	var product Product
	for rows.Next() {
		err = rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.ProductDesc,
			&product.ProductImage,
			&product.ProductPrice,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return Product{}, err
		}
		rowIndex++
	}

	if (Product{}) == product {
		msg := fmt.Sprintf("content with id %s not found", id)
		return Product{}, errors.New(msg)
	}

	return product, nil
}

// DeleteProductByID to delete
func (m *MDL) DeleteProductByID(id string) error {
	db := m.MysqlClient

	qs := `
	DELETE FROM product WHERE product_id = ?`

	_, err := db.Exec(qs, id)
	if err != nil {
		return err
	}

	return nil
}

// InsertProduct to insert product
func (m *MDL) InsertProduct(product Product) (int64, error) {
	db := m.MysqlClient

	qs := `
	INSERT INTO product ( product_id, product_name, product_desc, product_image, product_price) 
	VALUES (?, ?, ?, ?, ?);`

	result, err := db.Exec(qs, product.ProductID, product.ProductName, product.ProductDesc, product.ProductImage, product.ProductPrice)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// UpdateProductByID to update product
func (m *MDL) UpdateProductByID(id string, product Product) (int64, error) {
	db := m.MysqlClient

	qs := `
	UPDATE product SET product_name = ?, product_desc = ?, updated_at = ?, product_image = ?, product_price = ?
	WHERE product_id = ?;`

	result, err := db.Exec(qs, product.ProductName, product.ProductDesc, product.UpdatedAt, product.ProductImage, product.ProductPrice, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
