package main

import (
	"github.com/gin-gonic/gin"
)

func getSecretKey() string {
	return "UP5t13lSTvG97QOVPkmPkhsmhbI8qW"
}

func showRegistrationPage(c *gin.Context) {
	//render(c, gin.H{"title": "Register"}, "something.html")
}

func register(c *gin.Context) {
	addUser(c)
}
