package mock

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var SampleSecretKey = "MockSecretKey"

func GenerateMockToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        json.Number(strconv.FormatInt(time.Now().Add(5*time.Minute).Unix(), 10)),
		"authorized": true,
		"userId":     UserId,
	})
	tokenString, err := token.SignedString([]byte(SampleSecretKey))
	if err != nil {
		log.Fatalf("error generating token: %v", err)
	}

	return tokenString
}
