package output

import (
	"context"

	"github.com/marco-fpereira/to-do-list-server/adapters/exception"
	"github.com/marco-fpereira/to-do-list-server/config/logger"
	"github.com/marco-fpereira/to-do-list-server/domain/port/output"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type BCryptCryptographyAdapter struct{}

func NewBCryptCryptographyAdapter() output.CryptographyPort {
	return &BCryptCryptographyAdapter{}
}

func (bc *BCryptCryptographyAdapter) EncryptKey(rawKey string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawKey), 5)
	if err != nil {
		tag := []zap.Field{zap.String("cause", err.Error())}
		logger.Error(context.Background(), "Error encrypting key", err, tag...)
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
		tag := []zap.Field{zap.String("cause", err.Error())}
		logger.Error(context.Background(), "Error verifying encrypted key", err, tag...)
		return false
	}
	return true
}
