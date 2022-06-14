package utility

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pass string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	CheckError(err)
	return string(hashed)
}

func IsPasswordMatch(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return !CheckError(err)
}
