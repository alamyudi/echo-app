package dbmanager

import (
	"time"

	"github.com/alamyudi/echo-app/echokit/models"
)

type (
	// ProductManager model manager
	ProductManager struct {
		ProductID    string    `json:"product_id"`
		ProductName  string    `json:"product_name"`
		ProductDesc  string    `json:"product_desc"`
		ProductImage string    `json:"product_image"`
		ProductPrice float64   `json:"product_price"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

/********************* Product ***********************/

// GetProducts to getting products
func (m *DBManager) GetProducts() ([]models.Product, error) {
	products, err := m.MDL.GetProduct()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

// GetProductByID to getting product by id
func (m *DBManager) GetProductByID(id string) (models.Product, error) {
	product, err := m.MDL.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// DeleteProductByID to getting product by id
func (m *DBManager) DeleteProductByID(id string) error {
	err := m.MDL.DeleteProductByID(id)
	return err
}

// UpdateProductByID to update product by id
func (m *DBManager) UpdateProductByID(id string, product ProductManager) (int64, error) {
	productModel := models.Product{
		ProductName:  product.ProductName,
		ProductDesc:  product.ProductDesc,
		ProductImage: product.ProductImage,
		ProductPrice: product.ProductPrice,
		UpdatedAt:    product.UpdatedAt,
	}
	row, err := m.MDL.UpdateProductByID(id, productModel)
	return row, err
}

// InsertProduct to add product
func (m *DBManager) InsertProduct(product ProductManager) (int64, error) {
	productModel := models.Product{
		ProductID:    product.ProductID,
		ProductName:  product.ProductName,
		ProductDesc:  product.ProductDesc,
		ProductImage: product.ProductImage,
		ProductPrice: product.ProductPrice,
	}
	row, err := m.MDL.InsertProduct(productModel)
	return row, err
}
