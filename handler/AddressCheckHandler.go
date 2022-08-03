package handler

import (
	"net/http"

	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/Grama-Check/Address-Check-Api/util"
	"github.com/gin-gonic/gin"
)

func AddressCheck(c *gin.Context) {
	person := models.Person{}

	err := c.BindJSON(&person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.JsonError("Couldn't parse request body to json"))
	}

	
}
