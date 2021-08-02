package main

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func addUser(c *gin.Context){
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR] ":err})
		return
	}
	id, err := addDBUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"[Successfully created a new account]":id})
}


func addDBUser(user *User)(primitive.ObjectID, error){
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


func findUser(c *gin.Context)() {
	var user User
	if err := c.BindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR2]":err})
		return
	}
	var data, err = findDBUser(user.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"[ERROR3]": err})
		return 
	}
	c.JSON(http.StatusOK, 
		gin.H{
			"Id":data.Id, 
			"Email":data.Email, 
			"Username":data.Username, 
			"Password":data.Password, 
		})
}


func findDBUser(id primitive.ObjectID)(*User, error){
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
