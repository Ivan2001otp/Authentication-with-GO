package controllers

import (
	"context"
	"golang-org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strrconv"
	"time"

	"github.com/Ivan2001otp/Authentication-with-GO/helpers"
	"github.com/Ivan2001otp/Authentication-with-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/Ivan2001otp/Authentication-with-GO/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
);


var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user");
var validate = validator.New();


func HashPassword()


func SignUp()

func Login()

func VerifyPassword()
