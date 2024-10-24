package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func HashPasswordSHA256(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashEmailSHA256(email string) string {
	data := fmt.Sprintf("%s:%d", email, time.Now().UnixNano())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
