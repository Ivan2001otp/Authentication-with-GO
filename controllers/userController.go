package controllers

import (
	"context"
	"golang-org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strrconv"
	"time"

	"github.com/Ivan2001otp/Authentication-with-GO/database"
	"github.com/Ivan2001otp/Authentication-with-GO/helpers"
	"github.com/Ivan2001otp/Authentication-with-GO/models"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
);


var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user");
var validate = validator.New();


func HashPassword()


func SignUp() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second);
		var user models.User;

		if err := c.BindJSON(&user);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()});
			return;
		}

		validationErr := validate.Struct(user);

		if validationErr!=nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":validationErr.Error()});
			return;
		}

		count,err := userCollection.CountDocuments(ctx,bson.M{"email":user.Email});
		defer cancel();

		if err!= nil{
			log.Panic(err);
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Error obtained while checking email."});
			return;
		}

		count,err = userCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})

		defer cancel();

		if err!=nil{
			log.Panic(err);
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while checking for the phone number."});

		}

		if count>0{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"This email or phone number already exists."});
			
		}
	}
}

func Login()

func VerifyPassword()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id");

		if err := helpers.MatchUserTypetoUid(c,userId);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()});
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),100 * time.Second);

		var user models.User;

	   err :=	userCollection.FindOne(ctx,bson.M{"user_id":userId}).Decode(&user);
		
	   defer cancel();

	   if err!= nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()});
		return;
	}

	c.JSON(http.StatusOK,user);
	}
}