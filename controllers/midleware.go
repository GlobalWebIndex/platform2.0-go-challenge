package controllers

import (
	"gwi-challenge/common"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func removeBearer(token string) (string, error) {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:], nil
	}
	return token, nil
}

var authorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	removeBearer,
}

var auth2Extractor = &request.MultiExtractor{
	authorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", 0)
		token, err := request.ParseFromRequest(c.Request, auth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.GetConfig().Secret))
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["id"].(float64))
			c.Set("user_id", int(userID))
		}
	}
}
