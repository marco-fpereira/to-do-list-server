init-protobuf:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

init-dependency:
	go get -u github.com/antonfisher/nested-logrus-formatter
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/crypto
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql
	go get -u github.com/DATA-DOG/go-sqlmock
	go get -u github.com/go-crypt/crypt
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/golang-jwt/jwt/v5
	go get -u github.com/golang/mock/gomock
	go get -u github.com/google/uuid
	go get -u github.com/jarcoal/httpmock
	go get -u github.com/jinzhu/copier
	go get -u github.com/joho/godotenv
	go get -u github.com/sirupsen/logrus
	go get -u github.com/stretchr/testify/assert
	go get -u google.golang.org/grpc

test:
	go test ./tests/... -cover

build:
	go build -o bin/app app.go

run:
	go run app.go

deploy: init-protobuf init-dependency test build run
