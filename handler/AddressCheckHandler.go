package handler

import (
	"context"
	"database/sql"
	"net/http"

	db "github.com/Grama-Check/Address-Check-Api/db/sqlc"
	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://jhivan:25May2001@identitycheckserver.postgres.database.azure.com/postgres?sslmode=require"
)

func AddressCheck(c *gin.Context) {

	conn, err := sql.Open(dbDriver, dbSource)
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
