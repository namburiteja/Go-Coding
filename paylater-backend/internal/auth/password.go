package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword converts a plain-text password into a bcrypt hash.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword compares the hashed password stored in DB
// with the password entered by the user.
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}