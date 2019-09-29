package handlers

import (
	"gwiapi/app/models"
	"gwiapi/app/responses"
	"net/http"

	"github.com/jinzhu/gorm"
)

// GetAssets will bring asset of the user
func GetCountries(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	country := models.Country{}
	responses.JSON(w, http.StatusOK, country.GetAllCountries())
}
