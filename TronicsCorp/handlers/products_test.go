package handlers

import (
	"TronicsCorp/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	c *mongo.Client
	db *mongo.Database
	col *mongo.Collection
	cfg config.Properties
	h ProductHandler
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
func TestProduct(t *testing.T) {
	t.Run("TestProduct", func(t *testing.T) {
		body := `
		[{
			"product_name":"google pixel 3",
			"price":250,
			"currency":"USD",
			"vendor":"google",
			"accessories":["charger","subscription"]
		  }]
		`

		req := httptest.NewRequest("POST", "/products", strings.NewReader(body))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		h.Col = col
		err := h.CreateProducts(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusCreated, res.Code)
	})
	t.Run("get products", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		h.Col = col
		err := h.GetProducts(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusOK, res.Code)
	})
}