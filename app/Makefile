init-protobuf:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

init-dependency:
	go get github.com/antonfisher/nested-logrus-formatter
	go get github.com/gin-gonic/gin
	go get golang.org/x/crypto
	go get gorm.io/gorm
	go get gorm.io/driver/mysql
	go get github.com/DATA-DOG/go-sqlmock
	go get github.com/go-crypt/crypt
	go get github.com/go-sql-driver/mysql
	go get github.com/golang-jwt/jwt/v5
	go get github.com/golang/mock/gomock
	go get github.com/google/uuid
	go get github.com/jarcoal/httpmock
	go get github.com/jinzhu/copier
	go get github.com/joho/godotenv
	go get github.com/sirupsen/logrus
	go get github.com/stretchr/testify/assert
	go get go.uber.org/zap
	go get go.uber.org/zap/zapcore
	go get google.golang.org/grpc

test:
	go test ./tests/... -cover

build:
	go build -o bin/app app.go

run:
	go run app.go

deploy: init-protobuf init-dependency test build run
