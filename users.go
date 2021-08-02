package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func addUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR] ": err})
		return
	}

	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashed_password)

	id, err := addDBUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": err})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //expires after 1 day
	})

	// If the user is created, set the token in a cookie and log the user in
	token, err := claims.SignedString([]byte(getSecretKey()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
	}
	log.Println("token", token)
	c.SetCookie("token", token, 3600*24, "", "", true, true) //MaxAge: 24 hours
	c.Set("is_logged_in", true)

	c.JSON(http.StatusOK, gin.H{
		"[Successfully created a new account]": id,
		"Email":                                user.Email,
		"Username":                             user.Username,
		"Password":                             user.Password})
}

func addDBUser(user *User) (primitive.ObjectID, error) {
	client, ctx, cancel := connectToDatabase()
	defer cancel()
	defer client.Disconnect(ctx)
	user.Id = primitive.NewObjectID()
	data, err := client.Database("smartiez").Collection("users").InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("[ERROR] Could not create the account: %v", err)
		return primitive.NilObjectID, err
	}
	id := data.InsertedID.(primitive.ObjectID)
	return id, nil
}

func findUser(c *gin.Context) {
	var user User
	if err := c.BindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR2]": err})
		return
	}
	var data, err = findDBUser(user.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"[ERROR3]": err})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"Id":       data.Id,
			"Email":    data.Email,
			"Username": data.Username,
			"Password": data.Password,
		})
}

func findDBUser(id primitive.ObjectID) (*User, error) {
	var user *User
	client, ctx, cancel := connectToDatabase()
	defer cancel()
	defer client.Disconnect(ctx)
	data := client.Database("smartiez").Collection("users").FindOne(ctx, bson.D{})
	if data == nil {
		return nil, errors.New("[ERROR FOUND1]")
	}
	err := data.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("User: %v", user)
	return user, nil
}

/*
func findAllUsers(c *gin.Context) []*User {
	var users, err = findAllDBUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"[ERROR StatusNotFound]": err})
	}
	//c.JSON(http.StatusOK, gin.H{"available resources": users})
	log.Println("users: ", users)
	return users
}
func findAllDBUsers() ([]*User, error) {
	var users []*User
	client, ctx, cancel := connectToDatabase()
	defer cancel()
	defer client.Disconnect(ctx)
	data, err := client.Database("smartiez").Collection("users").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer data.Close(ctx)
	err = data.All(ctx, &users)
	if err != nil {
		fmt.Println("[ERROR 10]", err)
		return nil, err
	}
	return users, nil
}*/
