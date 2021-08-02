package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
}

type Resource struct {
	Id          primitive.ObjectID `json:"_id,omitempty"`
	Author      string             `json:"author,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Type        string             `json:"type,omitempty"`
	Date        string             `json:"date,omitempty"`
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func staticFile(c *gin.Context) {
	fileName := "my-app/build/" + c.Request.URL.Path
	c.File(fileName)
}

func main() {
	router := gin.Default()
	router.Use(setUserStatus())
	router.GET("/api/resource/:id", findResource)
	router.GET("/api/resource/", findAllResources)
	router.POST("/api/resource/", checkLoggedIn(), addResource)
	router.PUT("/api/resource/", checkLoggedIn(), updateResource)
	router.DELETE("/api/resource/:id", checkLoggedIn(), removeResource)
	router.GET("/api/users/:id", checkLoggedIn(), findUser)
	router.POST("/api/users/", checkLoggedIn(), addUser)
	// router.PUT("/api/users/",            updateUser)
	// router.DELETE("/api/users/:id",      removeUser)
	router.POST("/api/register", checkNotLoggedIn(), register)
	router.GET("/api/register", checkNotLoggedIn(), showRegistrationPage)
	router.POST("/api/login", checkNotLoggedIn(), login)
	router.GET("/api/login", checkNotLoggedIn(), showLoginPage)
	router.POST("/api/logout", checkLoggedIn(), logout)
	router.GET("/api/logout", checkLoggedIn(), showLogoutPage)
	router.NoRoute(staticFile)
	// router.PUT("/api/users/", 			updateUser)
	// router.DELETE("/api/users/:id", 		removeUser)
	router.Run()
}

func connectToDatabase() (*mongo.Client, context.Context, context.CancelFunc) {
	mongodb := goDotEnvVariable("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

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
	return client, ctx, cancel
}
