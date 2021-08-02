package main

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func addResource(c *gin.Context){
	var resource Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR] ":err})
		return
	}
	id, err := addDBResource(&resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"[Successfully created a new resource]":id})
}


func addDBResource(resource *Resource)(primitive.ObjectID, error){
	client, ctx, cancel := connectToDatabase()
	defer cancel()
	defer client.Disconnect(ctx)
	resource.Id = primitive.NewObjectID()
	data, err := client.Database("smartiez").Collection("resources").InsertOne(ctx, resource)
	if err != nil {
		fmt.Printf("[ERROR] Could not create the resource: %v", err)
		return primitive.NilObjectID, err
	}
	id := data.InsertedID.(primitive.ObjectID)
	return id, nil
}


func findResource(c *gin.Context)() {
	var resource Resource
	if err := c.BindUri(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR2]":err})
		return
	}
	var data, err = findDBResource(resource.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"[ERROR3]": err})
		return 
	}
	c.JSON(http.StatusOK, 
		gin.H{
			"Id":data.Id, 
			"Author":data.Author, 
			"Title":data.Title, 
			"Description":data.Description, 
			"Type":data.Type,
			"Date":data.Date,
		})
}


func findDBResource(id primitive.ObjectID)(*Resource, error){
	var resource *Resource
	client, ctx, cancel := connectToDatabase()
	defer cancel()
	defer client.Disconnect(ctx)
	data := client.Database("smartiez").Collection("resources").FindOne(ctx, bson.D{})
	if data == nil {
		return nil, errors.New("[ERROR FOUND1]")
	}
	err := data.Decode(&resource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Resource: %v", resource)
	return resource, nil
}


func findAllResources(c *gin.Context)(){
	var resources, err = findAllDBResources()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"[ERROR FOUND5]":err})
	}
	c.JSON(http.StatusOK, gin.H{"available resources":resources})
}


func findAllDBResources()([]*Resource, error) {
	var resources []*Resource
	client, ctx, cancel := connectToDatabase()
	defer cancel()
	defer client.Disconnect(ctx)
	data, err := client.Database("smartiez").Collection("resources").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer data.Close(ctx)
	err = data.All(ctx, &resources)
	if err != nil {
		fmt.Println("[ERROR 8]", err)
		return nil, err
	}
	return resources, nil
}


func updateResource(c *gin.Context) {

}


func removeResource(c *gin.Context) {

}
