package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product describes an eletronic product e.g. phone
type Product struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"product_name" bson:"product_name" validate:"required, max=10"`
	Price int `json:"price" bson:"price" validate:"required, max=2000"`
	Currency string `json:"currency" bson:"currency" validate:"required, len=3"`
	Quantity int `json:"quantity" bson:"quantity"`
	Discount int `json:"discount" bson:"discount"`
	Vendor string `json:"vendor" bson:"vendor" validate:"required"`
	Accessories []string `json:"accessories,omitempty" bson:"accessories,omitempty"`
	IsEssential bool `json:"is_essential" bson:"is_essential"`
}


// CreateProducts creates a new product on mongoDB
func CreateProducts(c echo.Context) error  {
	return c.JSON(http.StatusCreated, "Created")
}
	
