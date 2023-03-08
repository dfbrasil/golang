package tronics

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

//Start starts aplication

func Start(){
	port := os.Getenv("MY_APP_PORT")

	if port == ""{
		port = "8080"
	}

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.PUT("/products/:id", updateProduct)
	e.POST("/products", createProduct)

	e.StdLogger.Printf(fmt.Sprintf("Rodando echo na porta %s", port))
	e.Logger.Fatal(e.Start((fmt.Sprintf("localhost:%s", port))))
}