package tronics

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var v = validator.New()

func init(){
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%v+", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to laod configuration")
	}
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc{ //serverMessage tem a mesma assinatura de um middleware function padrão: aceita um handlerfunction como argumento e retorna um handlefunction. É um tipo que tem ma function que aceita uma handlerfunc e reuotna uma handlerfunc
	return func(c echo.Context) error { //serverMessage retorna uma function que contem uma mesma assinatura de um handler function. Também é um tipo que é uma função que recebe um contexto e retorna um error
		fmt.Println("inside serverheader middleware")
		return next(c)
	}
}

//Start starts aplication
func Start(){
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		e.Logger.Fatal("Unable to laod configuration")
	}
	e.Use(serverMessage)
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.PUT("/products/:id", updateProduct)
	e.POST("/products", createProduct)

	e.StdLogger.Printf(fmt.Sprintf("Rodando echo na porta %s", cfg.Port))
	e.Logger.Fatal(e.Start((fmt.Sprintf("localhost:%s", cfg.Port))))
}