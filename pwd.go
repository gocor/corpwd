package corpwd

import (
	"golang.org/x/crypto/bcrypt"
)

const pwdCost = 14

// HashPassword takes a users password and returns it hashed for storage
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), pwdCost)
	return string(bytes), err
}

// VerifyHashedPassword is valid
func VerifyHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
