package output

import (
	"context"
	"errors"
	"os"

	"github.com/marco-fpereira/to-do-list-server/adapters/exception"
	"github.com/marco-fpereira/to-do-list-server/config/logger"
	"github.com/marco-fpereira/to-do-list-server/domain/port/output"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
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
	accessToken, err := j.getJwtToken(token)
	if err != nil {
		return false, err
	}

	if !accessToken.Valid {
		return false, exception.BuildJwtException(401, "invalid token")
	}

	return true, nil
}

func (j *JwtAuthenticationAdapter) ValidateClaim(
	token string,
	claimName string,
	claimValue string,
) (bool, error) {
	accessToken, err := j.getJwtToken(token)
	if err != nil {
		return false, err
	}

	if jwtClaims, ok := accessToken.Claims.(jwt.MapClaims); ok {
		claims := make(map[string]string)
		for k, v := range jwtClaims {
			if strV, ok := v.(string); ok {
				claims[k] = string(strV)
			}
		}
		claim, ok := claims[claimName]

		if !ok {
			return false, exception.BuildJwtException(401, "unable to get claims to validate authenticity")
		}
		if claim != claimValue {
			return false, exception.BuildJwtException(403, "user does not contain required claims")
		}

		return true, nil
	}
	return false, exception.BuildJwtException(401, "unable to extract claims")
}

func (j *JwtAuthenticationAdapter) getJwtToken(token string) (*jwt.Token, error) {
	if token == "" {
		return nil, exception.BuildJwtException(401, "token not received")
	}
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("there's an error with the signing method")
			tag := []zap.Field{zap.String("cause", err.Error())}
			logger.Error(context.Background(), "error parsing token", err, tag...)

			return false, exception.BuildJwtException(401, "error parsing token")
		}
		return j.secretKey, nil
	})
}
