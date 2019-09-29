package handlers

import (
	"net/http"
	"time"

	"gwiapi/app/responses"

	"github.com/jinzhu/gorm"
)

// Ping will respond with Time of Call
func Ping(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, time.Now())
}
