package main

import (
	"github.com/Grama-Check/Address-Check-Api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authGroup := r.Group("/")

	authGroup.POST("/", handler.AddressCheck)

	r.Run("7070")
}
