package password

import (
	"crypto/sha256"
	"encoding/hex"
)

const PASS_SALT = "asdhgd3g478y bsfdfgiuy43wrw8nyfhy7384ytr"

func HashPassword(password string) string {
	data := password + PASS_SALT
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func ComparePassword(plainPassword, hashedPassword string) bool {
	return HashPassword(plainPassword) == hashedPassword
}
