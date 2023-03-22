package handlers

import (
	"TronicsCorp/dbiface"
	"context"
	"encoding/json"
	"io"

	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	v = validator.New()
)

// Product describes an eletronic product e.g. phone
type Product struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"product_name" bson:"product_name" validate:"required,max=40"`
	Price int `json:"price" bson:"price" validate:"required,max=2000"`
	Currency string `json:"currency" bson:"currency" validate:"required,len=3"`
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

// ProductValidator is a struct that contains the validator
type ProductValidator struct {
	validator *validator.Validate
}

// Validate validates the product
func (p *ProductValidator) Validate(i interface{}) error  {
	return p.validator.Struct(i)
}

func findProducts(ctx context.Context,q url.Values, collection dbiface.CollectionAPI) ([]Product, error)  {
	var products []Product
	filter := make(map[string]interface{})
	for key, value := range q {
		filter[key] = value[0]
	}
	if filter["_id"] != nil {
		docID, err := primitive.ObjectIDFromHex(filter["_id"].(string))
		if err != nil {
			return products, err
		}
		filter["_id"] = docID
	}
	cursor, err := collection.Find(ctx, bson.M(filter))
	if err != nil {
		log.Errorf("Unable to find products: %v", err)
		return products, err
	}
	err = cursor.All(ctx, &products)
	if err != nil {
		log.Errorf("Unable to decode products: %v", err)
		return products, err
	}
	return products, nil
}

// GerProduct returns a list of products
func (h *ProductHandler) GetProducts(c echo.Context) error  {
	products, err := findProducts(context.Background(),c.QueryParams(),  h.Col)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func insertProducts(ctx context.Context, products []Product, collection dbiface.CollectionAPI) ([]interface{}, error)  {
	var insertedIds []interface{}
	for _, product := range products {
		product.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Errorf("Unable to insert product: %v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

// CreateProducts creates a new product on mongoDB
func (h *ProductHandler) CreateProducts(c echo.Context) error  {
	var products []Product
	c.Echo().Validator = &ProductValidator{validator: v}
	if err := c.Bind(&products); err != nil {
		log.Errorf("Can not bind the request body to the product struct: %v", err)
		return err
	}
	for _, product := range products {
		if err := c.Validate(product); err != nil {
			log.Errorf("Can not validate the product: %+v %v" , product, err)
			return err
		}
	}
	IDs, err := insertProducts(context.Background(), products, h.Col)
	if err != nil {
		return err
	}
	log.Printf("Inserted products with IDs: %v", IDs)
	return c.JSON(http.StatusCreated, IDs)
}

// modifyProduct modifies a product
func modifyProduct(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (Product, error) {
	var product Product

	//find if the product exists, if err return 404
	docId, err := primitive.ObjectIDFromHex(id)
	log.Errorf("Unable to find product: %v", err)
	if err != nil {
		return product, err
	}
	filter := bson.M{"_id": docId}
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(&product); err != nil {
		log.Errorf("Unable to decode product: %v", err)
		return product, echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	//decode the request body to product struct, if err return 500
	if err := json.NewDecoder(reqBody).Decode(&product); err != nil {
		log.Errorf("Unable to decode the request body to product struct: %v", err)
		return product, echo.NewHTTPError(http.StatusInternalServerError, "Unable to decode the request body to product struct")
	}

	//validate the request body, if err return 400
	if err := v.Struct(product); err != nil {
		log.Errorf("Unable to validate the product: %v", err)
		return product, echo.NewHTTPError(http.StatusBadRequest, "Unable to validate the product")
	}

	//update the product, if err return 500
	_,err = collection.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		log.Errorf("Unable to update product: %v", err)
		return product, echo.NewHTTPError(http.StatusInternalServerError, "Unable to update product")
	}
	return product, nil
}

// UpdateProduct updates a product
func (h *ProductHandler) UpdateProduct(c echo.Context) error{
	product, err := modifyProduct(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		log.Errorf("Unable to update product: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, product)
}

// findProduct finds a product
func findProduct(ctx context.Context, id string, collection dbiface.CollectionAPI) (Product, error){
	var product Product
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}
	res := collection.FindOne(ctx, bson.M{"_id": docID})
	err = res.Decode(&product)
	if err != nil{
		return product, err
	}
	return product, nil
}

// GetProduct returns a product
func (h *ProductHandler) GetProduct(c echo.Context) error {
	product, err := findProduct(context.Background(), c.Param("id"), h.Col)
	if err != nil{
		return err
	}
	return c.JSON(http.StatusOK, product)
}

// deleteProduct deletes a product
func deleteProduct(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64,error){
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Unable to converto to ObjectID: %v", err)
		return 0, err
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		return 0, nil
	}
	return res.DeletedCount, nil
	
}// DeleteProduct deletes a product
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	delCount, err := deleteProduct(context.Background(), c.Param("id"), h.Col)
	if err != nil{
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}