package tests

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/marco-fpereira/to-do-list-server/tests/mock"

	jwt "github.com/golang-jwt/jwt/v5"
)

var SampleSecretKey = "MockSecretKey"

func GenerateMockToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        json.Number(strconv.FormatInt(time.Now().Add(5*time.Minute).Unix(), 10)),
		"authorized": true,
		"userId":     mock.UserId,
	})
	tokenString, err := token.SignedString([]byte(SampleSecretKey))
	if err != nil {
		log.Fatalf("error generating token: %v", err)
	}

	return tokenString
}

func SetEnvVars() {
	os.Setenv("HOST", "localhost")
	os.Setenv("DBPORT", "3306")
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "root")
	os.Setenv("DBNAME", "TODOLIST")
	os.Setenv("JWT-SECRET-KEY", SampleSecretKey)
}

func DeleteEnvVars() {
	os.Unsetenv("HOST")
	os.Unsetenv("DBPORT")
	os.Unsetenv("DBUSER")
	os.Unsetenv("DBPASS")
	os.Unsetenv("DBNAME")
}
