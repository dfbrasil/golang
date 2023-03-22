package handlers

import (
	"TronicsCorp/dbiface"
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User describes a user
type User struct{
	Email string `json:"username" bson:"username" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8,max=300"`
}

// UserHandler is a struct that contains the collection
type UsersHandler struct{
	Col dbiface.CollectionAPI
}

// UserValidator is a struct that contains the validator
type userValidator struct {
	validator *validator.Validate
}

// Validate validates the user
func (u *userValidator) Validate(i interface{}) error  {
	return u.validator.Struct(i)
}

// insertUser inserts a user into the database
func insertUser(ctx context.Context, user User, collection dbiface.CollectionAPI) (interface{}, *echo.HTTPError){
	var newUser User
	res := collection.FindOne(ctx, bson.M{"username": user.Email})
	err := res.Decode(&newUser)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("Unable to decode retrieved user: %v", err)
		return nil, echo.NewHTTPError(500, "Unable to decode retrieved user.")
	}
	if newUser.Email != "" {
		log.Errorf("User by %s already exists",user.Email)
		return nil, echo.NewHTTPError(400, "User already exists.")
	}
	insertRes, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Errorf("Unable to insert user: %v", err)
		return nil, echo.NewHTTPError(500, "Unable to insert user.")
	}
	return insertRes.InsertedID, nil
}

// CreateUser creates a user
func (h *UsersHandler) CreateUser(c echo.Context) error {
	var user User
	c.Echo().Validator = &userValidator{validator: v}
	if err := c.Bind(&user); err != nil {
		log.Errorf("Unable to bind to user struct.")
		return echo.NewHTTPError(400, "Unable to parse the request payload.")
	}
	if err := c.Validate(user); err != nil {
		log.Errorf("Unable to validate the reqeust body.")
		return echo.NewHTTPError(400, "Unable to validate the request payload.")
	}
	insertedUserID, err := insertUser(context.Background(), user, h.Col)
	if err != nil {
		log.Errorf("unable to insert to database")
		return err
	}
	return c.JSON(http.StatusCreated, insertedUserID)
}