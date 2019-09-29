package handlers

import (
	"fmt"
	"net/http"

	"gwiapi/app/models"
	"gwiapi/app/responses"

	"github.com/jinzhu/gorm"
)

func GetParticipants(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	participant := models.Participant{}

	participants, err := participant.GetAllParticipants(db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, participants)
}

func AddParticipant(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	participant := models.Participant{}

	err := db.Model(&models.Participant{}).Save(&participant).Error
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, participant.ID))
	responses.JSON(w, http.StatusOK, participant)
}
