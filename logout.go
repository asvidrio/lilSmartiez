package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showLogoutPage(c *gin.Context) {
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/api/resource/")
}
