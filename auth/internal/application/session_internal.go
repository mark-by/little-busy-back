package application

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"strconv"
)

func (s Session) createSessionValue(userID int) string {
	randomString := uuid.New().String()
	hash := sha256.Sum256([]byte(randomString + strconv.Itoa(userID)))
	return hex.EncodeToString(hash[:])
}
