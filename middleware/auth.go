package middleware

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Grama-Check/Address-Check-Api/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleWare() gin.HandlerFunc {
	buffer, err := ioutil.ReadFile("public.pem")
	if err != nil {
		log.Fatal("Cannot read publickey")
	}

	block, _ := pem.Decode(buffer)

	if err != nil {
		log.Fatal("Cannot parse public.pem file")
	}
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		log.Fatal("Cannot parse public key")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.JsonError("No token present"))
		}

		fields := strings.Fields(token)

		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.JsonError("Single string no prefix"))
		}

		if strings.ToLower(fields[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.JsonError("Incorrect prefix"))

		}
		jwtString := fields[1]
		claims := &jwt.RegisteredClaims{}
		parsedToken, err := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {
			return rsaPublicKey, nil
		})
		if !parsedToken.Valid || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()

	}
}
