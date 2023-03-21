package main

import (
	"TronicsCorp/config"
	"TronicsCorp/handlers"
	"context"
	"fmt"


	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CorrelationID is the name of the header that contains the correlation id
const(
	CorrelationID = "X-Correlation-ID"
)

var (
	c *mongo.Client
	db *mongo.Database
	col *mongo.Collection
	cfg config.Properties
)

func init()  {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration can not be read: %v", err)
	}
	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if  err != nil {
		log.Fatalf("Can not connect to MongoDB: %v", err)
	}
	db := c.Database(cfg.DBName)
	col = db.Collection(cfg.CollectionName)
}

func addCorrelationID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// generate a new correlation id if not present
		id := c.Request().Header.Get(CorrelationID)
		var newID string
		if id == "" {
			// generate reandom number
			newID = random.String(12)
			
		}else {
			newID = id
		}
		c.Request().Header.Set(CorrelationID, newID)
		c.Response().Header().Set(CorrelationID, newID)
		return next(c)
	}
}

func main()  {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(addCorrelationID)//Custon middleware
	h := &handlers.ProductHandler{Col: col}
	e.POST("/products", h.CreateProducts, middleware.BodyLimit("1M"))
	e.GET("/products", h.GetProducts)
	e.PUT("/products/:id", h.UpdateProduct, middleware.BodyLimit("1M"))
	e.Logger.Infof("Server is running on %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
