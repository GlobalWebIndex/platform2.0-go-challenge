package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func NotTestHash(t *testing.T) {
	hashedPass, _ := Hash("secretpassword")
	if string(hashedPass) != "$2a$10$tiQqO4rPzt8ViURj7aooV..kTFi/jS6EyuOBGcOkgWjRTz5v5Nfnm" {
		t.Errorf("Hash is not working, given %s and got %s and expected %s", "secretpassword", "$2a$10$tiQqO4rPzt8ViURj7aooV..kTFi/jS6EyuOBGcOkgWjRTz5v5Nfnm", string(hashedPass))
	}
}

func TestVerifyPassword(t *testing.T) {
	hashedPassword, _ := Hash("secretpassword")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("secretpassword"))
	if err != nil {
		t.Errorf("Comparing Hashed password does not work")
	}
}
