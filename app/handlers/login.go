package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gwiapi/app/auth"
	"gwiapi/app/models"
	"gwiapi/app/responses"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Login Check Credentials and if successfull respond with tokens
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(db, user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

// SignIn Helper func to Check Password
func SignIn(db *gorm.DB, email, password string) (string, error) {

	var err error

	userFound := models.User{}

	err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&userFound).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(userFound.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(userFound.ID)
}
