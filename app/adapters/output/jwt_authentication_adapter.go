package output

import (
	"errors"
	"os"
	"to-do-list-server/app/adapters/exception"
	"to-do-list-server/app/domain/port/output"

	jwt "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

type JwtAuthenticationAdapter struct {
	secretKey []byte
}

func NewJwtAuthenticationAdapter() output.AuthenticationPort {
	return &JwtAuthenticationAdapter{
		secretKey: []byte(os.Getenv("JWT-SECRET-KEY")),
	}
}

func (j *JwtAuthenticationAdapter) ValidateToken(
	token string,
) (bool, error) {
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, exception.BuildJwtException(errors.New("there's an error with the signing method"))
		}
		return j.secretKey, nil
	})

	if err != nil {
		log.WithField("error", err.Error()).Error("error parsing token")
		return false, exception.BuildJwtException(errors.New("error parsing token"))
	}

	if !accessToken.Valid {
		return false, exception.BuildJwtException(errors.New("invalid token"))
	}

	return true, nil
}
