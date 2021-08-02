package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "login unauthorized"})
			//c.JSON(http.StatusUnauthorized, gin.H{"status": "login unauthorized"})
		}
	}
}

func checkNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")

		loggedIn := loggedInInterface.(bool)
		log.Println("THis is the output for is_logged_in: ", loggedIn)
		if loggedIn {
			//c.AbortWithStatus(http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "already logged in, access to page not authorized"})
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Entererd setUserStatus")
		if token, err := c.Cookie("token"); err == nil || token != "" {
			log.Println("user is logged in ")
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
			log.Println("user is NOT logged in ")
		}
	}
}
