package tronics

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

//ProductValidator echo validatro for product
type ProductValidator struct{
	validator *validator.Validate
}

//Validate validates product request body
func (p *ProductValidator)  Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1:"mobiles"},{2:"tv"},{3:"laptops"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {

	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}		
	for _, p := range products{
		for k := range p{
			
			if pID == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {

	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}		
	for i, p := range products{
		for k := range p{
			if pID == k {
				product = p
				index = i
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	splice := func (s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
		//[1,2,3,4,5,6]
		//[1,2]+[4,5,6] = [1,2,4,5,6] (idiomatic in go)
	}

	products = splice(products, index)

	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return err
	}

	for _, p := range products{
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}

	if product == nil{
		return c.JSON(http.StatusNotFound, "product not found")
	}

	type body struct{
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body

	e.Validator = &ProductValidator{validator: v}

	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil{
		return err
	}

	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, product)
	}

func createProduct(c echo.Context) error {
	type body struct{
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body

	e.Validator = &ProductValidator{validator: v}

	if err := c.Bind(&reqBody); err != nil {
			return err
	}
	if err := c.Validate(reqBody); err != nil{
		return err
	}

	product := map[int]string{
		len(products) + 1: reqBody.Name, //cria um novo produto usando o product_name da request e ID do Map
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}
