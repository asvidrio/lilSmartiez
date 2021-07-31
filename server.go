package main

import (
	"github.com/gin-gonic/gin"
)

var PORT = ":9000"

func main() {
	r := gin.Default()
	r.Run(PORT)
}