package services

import (
	"gwi-challenge/data/models"
	"testing"
)

func TestEncryption(t *testing.T) {
	user := &models.User{
		Username:     "username",
		PasswordHash: "password123",
		FullName:     "fullname",
	}
	password := user.PasswordHash
	err := encryptPassword(user)
	if password == user.PasswordHash || err != nil {
		t.Errorf("password was not ecrypted")
	}

	user.PasswordHash = ""

	err = encryptPassword(user)
	shouldBeError(err, t)
}

func TestDecryption(t *testing.T) {

	user := &models.User{
		Username:     "username",
		PasswordHash: "password123",
		FullName:     "fullname",
	}

	validPassword := user.PasswordHash

	err := encryptPassword(user)
	shouldBeNull(err, t)

	err = checkPassword(user, "wrong password")
	shouldBeError(err, t)

	err = checkPassword(user, "")
	shouldBeError(err, t)

	err = checkPassword(user, validPassword+" ")
	shouldBeError(err, t)

	err = checkPassword(user, validPassword)
	shouldBeNull(err, t)
}

func shouldBeNull(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expected err to be nil but it was %s", err)
	}
}

func shouldBeError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("expected error but it was nil")
	}
}
