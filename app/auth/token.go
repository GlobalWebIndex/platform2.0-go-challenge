package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gwiapi/config"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken Generates a token
func CreateToken(uid uint32) (string, error) {
	config := config.GetConfig()
	expirationInMinutes := config.AUTH.ExpirationMinutes
	hmacSecret := config.AUTH.HmacSecret

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationInMinutes)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(hmacSecret))

}

// TokenValid Checks is Token provided is Valid
func TokenValid(r *http.Request) error {
	config := config.GetConfig()
	tokenString := ExtractToken(r)
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.AUTH.HmacSecret), nil
	})
	if err != nil {
		return err
	}
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	Pretty(claims)
	// }
	return nil
}

// ExtractToken Extracts "token" from Request Authorization Header
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID Extracts "uid" from Request authorization header
func ExtractTokenID(r *http.Request) (uint32, error) {

	config := config.GetConfig()
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AUTH.HmacSecret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["uid"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
