package handlers

import (
	"TronicsCorp/dbiface"
	"context"
	"log"

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
// ProductHandler is a struct that contains the collection
type ProductHandler struct{
	Col dbiface.CollectionAPI
}

func insertProducts(ctx context.Context, products []Product, collection dbiface.CollectionAPI) ([]interface{}, error)  {
	var insertedIds []interface{}
	for _, product := range products {
		product.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Printf("Unable to insert product: %v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

// CreateProducts creates a new product on mongoDB
func (h *ProductHandler) CreateProducts(c echo.Context) error  {
	var products []Product
	if err := c.Bind(&products); err != nil {
		log.Printf("Can not bind the request body to the product struct: %v", err)
		return err
	}
	IDs, err := insertProducts(context.Background(), products, h.Col)
	if err != nil {
		return err
	}
	log.Printf("Inserted products with IDs: %v", IDs)
	return c.JSON(http.StatusCreated, IDs)
}
	
