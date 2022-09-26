package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Message struct {
	Message string `json:"message"`
}

var (
	products = make([]*Product, 0)
	nomor    = 1
)

func CreateProduct(c echo.Context) error {
	p := &Product{
		ID: nomor,
	}

	msg := Message{}

	if err := c.Bind(p); err != nil {
		msg.Message = "Invalid Request Body"
		return c.JSON(http.StatusBadRequest, msg)
	}

	products = append(products, p)
	nomor++

	return c.JSON(http.StatusCreated, p)
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	paramID := c.Param("id")
	productID, err := strconv.Atoi(paramID)

	msg := Message{}

	if err != nil {
		msg.Message = "Invalid Request Parameter"
		return c.JSON(http.StatusBadRequest, msg)
	}

	var result *Product

	for _, p := range products {
		if p.ID == productID {
			if err := c.Bind(p); err != nil {
				msg.Message = "Invalid Request Body"
				return c.JSON(http.StatusBadRequest, msg)
			}
			result = p
		}
	}

	if result == nil {
		msg.Message = "Product Not Found"
		return c.JSON(http.StatusNotFound, msg)
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	paramID := c.Param("id")
	productID, err := strconv.Atoi(paramID)

	msg := Message{
		Message: "Delete Product Succeed",
	}

	if err != nil {
		msg.Message = "Invalid Request Parameter"
		return c.JSON(http.StatusBadRequest, msg)
	}

	var result []*Product

	for _, p := range products {
		if p.ID != productID {
			result = append(result, p)
		}
	}

	if len(products) == len(result) {
		msg.Message = "Product Not Found"
		return c.JSON(http.StatusNotFound, msg)
	}

	products = result

	return c.JSON(http.StatusOK, msg)
}

func main() {
	e := echo.New()

	e.POST("/", CreateProduct)
	e.GET("/", GetProducts)
	e.PUT("/:id", UpdateProduct)
	e.DELETE("/:id", DeleteProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
