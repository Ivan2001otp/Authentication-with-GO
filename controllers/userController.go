package controllers

import (
	"github.com/Ivan2001otp/Authentication-with-GO/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func SignUp()

func Login()

func VerifyPassword()
