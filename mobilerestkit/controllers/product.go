package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/alamyudi/echo-app/echokit/dbmanager"

	"github.com/labstack/echo"
)

type (

	// AddProductPayload add product payload
	AddProductPayload struct {
		ProductID    string  `validate:"required" json:"product_id"`
		ProductName  string  `validate:"required" json:"product_name"`
		ProductDesc  string  `validate:"required" json:"product_desc"`
		ProductImage string  `validate:"required" json:"product_image"`
		ProductPrice float64 `validate:"required" json:"product_price"`
	}

	// UpdateProductPayload add product payload
	UpdateProductPayload struct {
		ProductName  string  `validate:"required" json:"product_name"`
		ProductDesc  string  `validate:"required" json:"product_desc"`
		ProductImage string  `validate:"required" json:"product_image"`
		ProductPrice float64 `validate:"required" json:"product_price"`
	}
)

/******************* Product ********************/

// GetProducts to getting products
func GetProducts(ctx echo.Context) error {
	products, err := iKit.DBManager.GetProducts()
	if err != nil {
		logrus.Infof(err.Error())
		response := MessageResponse{
			Title:   "Failed getting products",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}
	response := MessageWithPayloadResponse{
		Title:   "Success",
		Message: "Success fetch products",
		Payload: products,
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetProductByID to getting products
func GetProductByID(ctx echo.Context) error {
	productID := ctx.Param("id")
	product, err := iKit.DBManager.GetProductByID(productID)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed getting product",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	response := MessageWithPayloadResponse{
		Title:   "Success",
		Message: "Success fetch product",
		Payload: product,
	}

	return ctx.JSON(http.StatusOK, response)
}

// DeleteProductByID to getting products
func DeleteProductByID(ctx echo.Context) error {
	productID := ctx.Param("id")
	err := iKit.DBManager.DeleteProductByID(productID)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed getting product",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	message := fmt.Sprintf("Success deleted product with id %s", productID)
	response := MessageResponse{
		Title:   "Success",
		Message: message,
	}

	return ctx.JSON(http.StatusOK, response)
}

// PutProductByID to getting products
func PutProductByID(ctx echo.Context) error {
	productID := ctx.Param("id")
	payload := new(UpdateProductPayload)

	// serialize json body
	if err := ctx.Bind(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Failed to serialize product",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	// validate interface
	if err := ctx.Validate(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	product := dbmanager.ProductManager{
		ProductName:  payload.ProductName,
		ProductDesc:  payload.ProductDesc,
		ProductImage: payload.ProductImage,
		ProductPrice: payload.ProductPrice,
		UpdatedAt:    time.Now(),
	}

	_, err := iKit.DBManager.UpdateProductByID(productID, product)
	if err != nil {
		logrus.Info(err.Error())
		response := MessageResponse{
			Title:   "Failed updating product",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	response := MessageResponse{
		Title:   "Success",
		Message: "Success updating product",
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostProduct to getting products
func PostProduct(ctx echo.Context) error {
	payload := new(AddProductPayload)

	// serialize json body
	if err := ctx.Bind(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Failed to serialize product",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	// validate interface
	if err := ctx.Validate(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	product := dbmanager.ProductManager{
		ProductID:    payload.ProductID,
		ProductName:  payload.ProductName,
		ProductDesc:  payload.ProductDesc,
		ProductImage: payload.ProductImage,
		ProductPrice: payload.ProductPrice,
		CreatedAt:    time.Now(),
	}

	_, err := iKit.DBManager.InsertProduct(product)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed adding product",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	response := MessageResponse{
		Title:   "Success",
		Message: "Success added product",
	}

	return ctx.JSON(http.StatusOK, response)
}
