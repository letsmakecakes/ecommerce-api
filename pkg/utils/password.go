package utils

import "golang.org/x/crypto/bcrypt"

// SetPassword hashes the password before storing it in the database.
func SetPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

// CheckPassword compares the hashed password with the plain-text password.
func CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password))
	return err == nil
}
