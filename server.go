package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var PORT = ":9000"

type User struct {
	Id primitive.ObjectID	`json:"_id,omitempty"`
	Email string		 	`json:"email,omitempty"`
	Username string			`json:"username,omitempty,unique"`
	Password string			`json:"password,omitempty"`
}

type Resource struct {
	Id primitive.ObjectID	`json:"_id,omitempty"`
	Author string 			`json:"author,omitempty"`
	Title string 			`json:"title,omitempty"`
	Description string 		`json:"description,omitempty"`
	Type string 			`json:"type,omitempty"`
	Date string 			`json:"date,omitempty"`
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	// routing logic
	router := gin.Default()
	router.GET("/api/resource/:id", 	findResource)
	router.GET("/api/resource/", 		findAllResources)
	router.POST("/api/resource/", 		addResource)
	router.PUT("/api/resource/", 		updateResource)
	router.DELETE("/api/resource/:id", 	removeResource)

	// router.GET("/api/users/:id", 		findUser)
	// router.GET("/api/users/", 			findAllUsers)
	// router.POST("/api/users/", 			addUser)
	// router.PUT("/api/users/", 			updateUser)
	// router.DELETE("/api/users/:id", 	removeUser)
	router.Run()
}

func connectToDatabase()(*mongo.Client, context.Context, context.CancelFunc) {
	// database logic
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

// 
// 
// 

func addUser(c *gin.Context){
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"[ERROR] ":err})
		return
	}
	id, err := CreateDBUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"[ERROR]": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"[Successfully created a new account]":id})
}

func CreateDBUser(user *User)(primitive.ObjectID, error){
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

func ReadDBUser(id primitive.ObjectID)(*User, error){
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

func ReadAllDBUsers()([]*User, error) {
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
		fmt.Println("[ERROR 8]", err)
		return nil, err
	}
	return users, nil
}