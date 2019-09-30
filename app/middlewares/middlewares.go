package middlewares

import (
	"errors"
	"net/http"

	"gwiapi/app/auth"
	"gwiapi/app/responses"

	"github.com/jinzhu/gorm"
)

// RequestHandlerFunction Overloading HandlerFunc with db
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

// JSON Adds Json Content Type
func JSON(next RequestHandlerFunction, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(db, w, r)
	}
}

// Auth Validates Tokens
func Auth(next RequestHandlerFunction, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		next(db, w, r)
	}
}
