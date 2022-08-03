package main

import (
	"github.com/Grama-Check/Address-Check-Api/handler"
	"github.com/Grama-Check/Address-Check-Api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authGroup := r.Group("/").Use(middleware.AuthMiddleWare())

	authGroup.POST("/", handler.AddressCheck)

	r.Run(":7070")
}
