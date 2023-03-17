package main

import (
	"TronicsCorp/config"
	"TronicsCorp/handlers"
	"context"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func main()  {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	h := &handlers.ProductHandler{Col: col}
	e.POST("/products", h.CreateProducts, middleware.BodyLimit("1M"))
	e.Logger.Infof("Server is running on %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
