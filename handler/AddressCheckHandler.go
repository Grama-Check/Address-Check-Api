package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	db "github.com/Grama-Check/Address-Check-Api/db/sqlc"
	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/Grama-Check/Address-Check-Api/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var config util.Config

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")

	}
}

func AddressCheck(c *gin.Context) {

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't connect to database")

		return
	}

	query := db.New(conn)

	person := models.Person{}

	err = c.BindJSON(&person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse request body to json")
		return
	}

	_, err = query.GetPerson(context.Background(), person.NIC)

	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"exists": false,
				"nic":    person.NIC,
			},
		)
	}

}
