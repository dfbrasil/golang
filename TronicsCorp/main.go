package main

import (
	"TronicsCorp/config"
	"TronicsCorp/handlers"
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/labstack/gommon/random"
	"go.mongodb.org/mongo-driver/bson"
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
	userCol *mongo.Collection
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
	col = db.Collection(cfg.ProductCollection)
	userCol = db.Collection(cfg.UsersCollection)

	isUserIndexUnique := true
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"username", 1}},
		Options: &options.IndexOptions{
			Unique: &isUserIndexUnique,
		},	
	}
	_, err = userCol.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Can not create index: %v", err)
	}
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
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(addCorrelationID)//Custon middleware
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.JwrTokenSecret),
		TokenLookup: "header:x-auth-token",
	})
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${remote_ip} ${header:X-Correlation-ID} ${host} ${method} ${uri} ${method} ${user_agent}` + ` ${status} ${error} ${latency_human}` + "\n",
	}))
	h := &handlers.ProductHandler{Col: col}
	uh := &handlers.UsersHandler{Col: userCol}
	e.POST("/products", h.CreateProducts, middleware.BodyLimit("1M"), jwtMiddleware)
	e.GET("/products", h.GetProducts)
	e.GET("/products/:id", h.GetProduct)
	e.DELETE("products/:id", h.DeleteProduct, jwtMiddleware)
	e.PUT("/products/:id", h.UpdateProduct, middleware.BodyLimit("1M"), jwtMiddleware)

	e.POST("/users", uh.CreateUser)
	e.POST("/auth", uh.AuthnUser)
	e.Logger.Infof("Server is running on %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
