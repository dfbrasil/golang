package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	port := os.Getenv("MY_APP_PORT")

	if port == ""{
		port = "8080"
	}

	e := echo.New()

	v := validator.New()

	products := []map[int]string{{1:"mobiles"},{2:"tv"},{3:"laptops"}} // slice de map com integer como key e string como value

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello there")
	})

	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	e.GET("/products/:id", func(c echo.Context) error {

		var product map[int]string

		for _, p := range products{
			for k := range p{
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct{
			Name string `json:"product_name" validate:"required,min=4"`
			// Vendor string `json:"vendor" validate: "min=5,max=10"`
			// Email string `json:"email" validate: "required_with=Vendor,email"`
			// Website string `json:"website" validate:"url"`
			// Country string `json:"country" validate: "len=2"`
			// DefaultDeviceIp string `json:"default_device_ip" validate:"ip"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil{
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name, //cria um novo produto usando o product_name da request e ID do Map
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})


	e.StdLogger.Printf(fmt.Sprintf("Rodando echo na porta %s", port))
	e.Logger.Fatal(e.Start((fmt.Sprintf("localhost:%s", port))))

	// e.StdLogger.Print("Rodando echo na porta 8080 STDLOGGER")
	// e.Logger.Fatal(e.Start(":8080"))

}