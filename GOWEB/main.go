package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main(){

	port := os.Getenv("MY_APP_PORT")

	if port == ""{
		port = "8080"
	}

	e := echo.New()

	products := []map[int]string{{1:"mobiles"},{2:"tv"},{3:"laptops"}} // slice de map com integer como key e string como value

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello there")
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

	e.StdLogger.Printf(fmt.Sprintf("Rodando echo na porta %s", port))
	e.Logger.Fatal(e.Start((fmt.Sprintf("localhost:%s", port))))

	// e.StdLogger.Print("Rodando echo na porta 8080 STDLOGGER")
	// e.Logger.Fatal(e.Start(":8080"))

}