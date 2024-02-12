package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
