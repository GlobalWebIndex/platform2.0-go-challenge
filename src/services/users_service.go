package services

import (
	"errors"
	"gwi-challenge/common"
	"gwi-challenge/data/models"
	"gwi-challenge/data/repositories"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// GetUserWithAccessToken returns user with its access token (login)
func GetUserWithAccessToken(username string, password string) (user *models.User, accessToken string, err error) {
	user, err = repositories.GetUserByUsername(username, false)
	if err != nil {
		return nil, "", errors.New("Invalid username or password")
	}

	err = checkPassword(user, password)
	if err != nil {
		return nil, "", errors.New("Invalid username or password")
	}

	accessToken = generateJWTToken(user.ID)
	return
}

// CreateNewUser creates a new user with a hashed password (register)
func CreateNewUser(userToCreate *models.User) error {
	encryptPassword(userToCreate)
	err := repositories.Save(userToCreate)
	if err != nil && strings.Contains(err.Error(), "constraint") {
		err = errors.New("username already exists")
	}
	return err
}

func GetUserById(userId uint, withFavorites bool) (userToReturn *models.User, err error) {
	userToReturn, err = repositories.GetUserByID(userId, withFavorites)
	return
}

func encryptPassword(user *models.User) error {
	if len(user.PasswordHash) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(user.PasswordHash)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.PasswordHash = string(passwordHash)
	return nil
}

func checkPassword(user *models.User, passwordToCheck string) error {
	bytePassword := []byte(passwordToCheck)
	byteHashedPassword := []byte(user.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func generateJWTToken(id uint) string {
	JwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	JwtToken.Claims = jwt.MapClaims{
		"id":     id,
		"expiry": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := JwtToken.SignedString([]byte(common.GetConfig().Secret))
	return token
}
