package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.GET("/users/:id", getUser)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func main(){

	e := echo.New()
	// e.GET("/", func (c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!!")
	// })

	e.Logger.Fatal(e.Start(":1323"))

}
