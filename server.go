package main

import (
	"log"
	"fmt"
	"context"
	"time"
	"os"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var PORT = ":9000"

var Users []struct {
	Id primitive.ObjectID
	Email string	
	Username string			
	Password string			
}

var Resource []struct {
	Id int
	Author string 
	Title string 
	Description string 
	Type string 
	Date string 
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	mongodb := goDotEnvVariable("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	users, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	resources, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users, resources)
}

// func addResource(c *gin.Context){
	
// }

// func findResource(c *gin.Context) {

// }

// func updateResource(c *gin.Context) {

// }

// func removeResource(c *gin.Context) {

// }