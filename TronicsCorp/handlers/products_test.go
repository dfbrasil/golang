package handlers

import (
	// "TronicsCorp/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	// c *mongo.Client
	db *mongo.Database
	col *mongo.Collection
	// cfg config.Properties
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
	col = db.Collection(cfg.ProductCollection)
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	testCode := m.Run()
	// Drop the test database
	col.Drop(ctx)
	// Close the connection
	db.Drop(ctx)
	os.Exit(testCode)
}

func TestProduct(t *testing.T) {
	var docID string
	t.Run("TestProduct", func(t *testing.T) {
		var IDs []string
		body := `
		[{
			"product_name":"google pixel 3",
			"price":250,
			"currency":"USD",
			"vendor":"google",
			"accessories":["charger","subscription"]
		  }]
		`

		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		h.Col = col
		err := h.CreateProducts(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusCreated, res.Code)

		err = json.Unmarshal(res.Body.Bytes(), &IDs)
		assert.Nil(t, err)
		docID = IDs[0]
		t.Logf("IDs: %#+v\n", IDs)
		for _, ID := range IDs {
			assert.NotNil(t, ID )
		}
	})
	t.Run("get products", func(t *testing.T) {
		var products []Product
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		h.Col = col
		err := h.GetProducts(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusOK, res.Code)

		err = json.Unmarshal(res.Body.Bytes(), &products)
		assert.Nil(t, err)
		for _, product := range products {
			assert.Equal(t, "google pixel 3", product.Name)
		}
	})

	t.Run("get product with query params", func(t *testing.T) {
		var products []Product
		req := httptest.NewRequest(http.MethodGet, "/products?vendor=google&price=250", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		h.Col = col
		err := h.GetProducts(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusOK, res.Code)

		err = json.Unmarshal(res.Body.Bytes(), &products)
		assert.Nil(t, err)
		for _, product := range products {
			assert.Equal(t, "google pixel 3", product.Name)
		}
	})

	t.Run("get a product", func(t *testing.T) {
		var product Product
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", docID),  nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		// c.SetParamValues(fmt.Sprintf("%s",docID))
		h.Col = col
		err := h.GetProduct(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusOK, res.Code)

		err = json.Unmarshal(res.Body.Bytes(), &product)
		assert.Nil(t, err)
		assert.Equal(t, "google pixel 3", product.Name)
	})

	t.Run("update a product", func(t *testing.T) {
		var product Product
		body := `
		{
			"product_name":"google pixel 5",
			"price":250,
			"currency":"USD",
			"vendor":"google",
			"accessories":["charger","subscription"]
		  }
		`

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%s", docID),  strings.NewReader(body))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		// c.SetParamValues(fmt.Sprintf("%s",docID))
		h.Col = col
		err := h.UpdateProduct(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusOK, res.Code)

		err = json.Unmarshal(res.Body.Bytes(), &product)
		assert.Nil(t, err)
		assert.Equal(t, "google pixel 5", product.Name)
	})

	t.Run("delete a product", func(t *testing.T) {
		var delCount int64
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%s", docID),  nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		// c.SetParamValues(fmt.Sprintf("%s",docID))
		h.Col = col
		err := h.DeleteProduct(c)
		assert.Nil(t, err )
		assert.Equal(t, http.StatusOK, res.Code)
		err = json.Unmarshal(res.Body.Bytes(), &delCount)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), delCount)
	})
}