package output

import (
	"to-do-list-server/app/adapters/exception"
	"to-do-list-server/app/domain/port/output"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type BCryptCryptographyAdapter struct{}

func NewBCryptCryptographyAdapter() output.CryptographyPort {
	return &BCryptCryptographyAdapter{}
}

func (bc *BCryptCryptographyAdapter) EncryptKey(rawKey string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawKey), 5)
	if err != nil {
		log.WithField("cause", err.Error()).Error("Error encrypting key")
		return "", exception.BuildBCryptException("Error encrypting key")
	}
	return string(bytes), nil
}

func (bc *BCryptCryptographyAdapter) VerifyEncryptedKey(
	rawKey string,
	encryptedKey string,
) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedKey), []byte(rawKey))
	if err != nil {
		log.WithField("cause", err.Error()).Error("Error verifying encrypted key")
		return false
	}
	return true
}
