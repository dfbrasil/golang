package handlers

import (
	"TronicsCorp/config"
	"TronicsCorp/dbiface"
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// User describes a user
type User struct{
	Email string `json:"username" bson:"username" validate:"required,email"`
	Password string `json:"password,omitempty" bson:"password" validate:"required,min=8,max=300"`
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

var (
	cfg config.Properties
)

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Errorf("Unable to hash password: %v", err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Unable to hash password.")
	}
	user.Password = string(hashedPassword)
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Errorf("Unable to insert user: %v", err)
		return nil, echo.NewHTTPError(500, "Unable to insert user.")
	}
	// insertRes, err := collection.InsertOne(ctx, user)
	// if err != nil {
	// 	log.Errorf("Unable to insert user: %v", err)
	// 	return nil, echo.NewHTTPError(500, "Unable to insert user.")
	// }
	return User{Email: user.Email}, nil
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

// AuthnUser authenticates a user
func (h *UsersHandler) AuthnUser(c echo.Context) error {
	var user User
	c.Echo().Validator = &userValidator{validator: v}
	if err := c.Bind(&user); err != nil {
		log.Errorf("Unable to bind to user struct.")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Unable to parse the request payload.")
	}
	if err := c.Validate(user); err != nil {
		log.Errorf("Unable to validate the reqeust body.")
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to validate the request payload.")
	}
	user, err := authenticateUser(context.Background(), user, h.Col)
	if err != nil {
		log.Errorf("unable to authenticate user")
		return err
	}
	token, er := createToken(user.Email)
	if er != nil {
		log.Errorf("unable to create token")
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to create token.")
	}
	c.Response().Header().Set("x-auth-token", token)
	return c.JSON(http.StatusOK, User{Email: user.Email})
}

// authenticateUser authenticates a user
func authenticateUser(ctx context.Context, reqUser User, collection dbiface.CollectionAPI) (User, *echo.HTTPError)  {
	var storedUser User //Useri n DB
	res := collection.FindOne(ctx, bson.M{"username": reqUser.Email})
	err := res.Decode(&storedUser)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("Unable to decode retrieved user: %v", err)
		return storedUser, echo.NewHTTPError(http.StatusUnprocessableEntity, "Unable to decode retrieved user.")
	}
	if err == mongo.ErrNoDocuments {
		log.Errorf("User by %s does not exist", reqUser.Email)
		return storedUser, echo.NewHTTPError(http.StatusNotFound, "User does not exist.")
	}

	// validate the password
	if !isCredValid(reqUser.Password, storedUser.Password) {
		log.Errorf("Incorrect password")
		return storedUser, echo.NewHTTPError(http.StatusUnauthorized, "Incorrect password.")
	}
	return User{Email: storedUser.Email}, nil
}

func isCredValid(givenPwr, storedPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(storedPwd), []byte(givenPwr)); err != nil{
		return false
	}
	return true
}

func createToken(username string) (string, error) {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Unable to read environment variables: %v", err)
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(cfg.JwrTokenSecret))
	if err != nil {
		log.Errorf("Unable to sign token: %v", err)
		return "", err
	}
	return token, nil
}
