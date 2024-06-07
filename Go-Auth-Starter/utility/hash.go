package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hashedPassword), err
}

// Returns nil if they match, or an error if they don't.
func CompareHashedPasswords(hashedPassword, password string) bool {
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
