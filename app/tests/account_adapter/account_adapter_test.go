package tests

import (
	"context"
	"os"
	"testing"
	"to-do-list-server/app/adapters/output"
	pb "to-do-list-server/app/config/grpc"
	tests "to-do-list-server/app/tests"
	mock "to-do-list-server/app/tests/mock"

	goSqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	tests.SetEnvVars()
	defer tests.DeleteEnvVars()
	mock.Ctx = context.Background()
	mock.Token = tests.GenerateMockToken()
	code := m.Run()
	os.Exit(code)
}

func TestSignup_200(t *testing.T) {
	assert := assert.New(t)
	client, sqlMock, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	sqlMock.ExpectQuery("SELECT * FROM `ACCOUNT` WHERE Username = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?").
		WithArgs(mock.Username, 1).WillReturnError(gorm.ErrRecordNotFound)

	sqlMock.ExpectBegin()

	sqlMock.ExpectExec(
		"INSERT INTO `ACCOUNT` (`UserId`,`Username`,`Password`) VALUES (?,?,?)",
	).WithArgs(mock.AnyString{}, mock.Username, mock.AnyString{}).WillReturnResult(
		goSqlMock.NewResult(1, 1),
	)

	sqlMock.ExpectCommit()

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  mock.Password,
		RequestId: uuid.New().String(),
		Token:     mock.Token,
	}

	void, err := client.Signup(mock.Ctx, request)
	if err != nil {
		t.Fatalf("error testing Signup: %v", err)
	}
	assert.NotNil(void)
}

func TestSignup_400UserAlreadyExist(t *testing.T) {
	assert := assert.New(t)
	client, sqlMock, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	sqlMock.ExpectQuery(
		"SELECT * FROM `ACCOUNT` WHERE Username = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?",
	).WithArgs(mock.Username, 1).WillReturnRows(
		sqlMock.NewRows([]string{"UserId", "Username", "Password"}).
			AddRow(mock.UserId, mock.Username, mock.Password),
	)

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  mock.Password,
		RequestId: uuid.New().String(),
		Token:     mock.Token,
	}

	void, err := client.Signup(mock.Ctx, request)
	assert.Equal("rpc error: code = Code(409) desc = err: user already exists | fields: [Username]", err.Error())
	assert.Nil(void)
}

func TestSignup_401TokenNotReceived(t *testing.T) {
	assert := assert.New(t)
	client, _, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  mock.Password,
		RequestId: uuid.New().String(),
	}

	void, err := client.Signup(mock.Ctx, request)
	assert.Equal("rpc error: code = Code(401) desc = err: token not received", err.Error())
	assert.Nil(void)
}

func TestSignup_400PasswordDoesNotMatchRequirements(t *testing.T) {
	assert := assert.New(t)
	client, sqlMock, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	sqlMock.ExpectQuery("SELECT * FROM `ACCOUNT` WHERE Username = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?").
		WithArgs(mock.Username, 1).WillReturnError(gorm.ErrRecordNotFound)

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  "Password123",
		RequestId: uuid.New().String(),
		Token:     mock.Token,
	}

	void, err := client.Signup(mock.Ctx, request)
	assert.Equal("rpc error: code = Code(400) desc = err: password is not strong enough | fields: []", err.Error())
	assert.Nil(void)
}

func TestLogin_200(t *testing.T) {
	assert := assert.New(t)
	client, sqlMock, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	bcrypt := output.NewBCryptCryptographyAdapter()
	encryptedPassword, _ := bcrypt.EncryptKey(mock.Password)

	sqlMock.ExpectQuery(
		"SELECT * FROM `ACCOUNT` WHERE Username = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?",
	).WithArgs(mock.Username, 1).WillReturnRows(
		sqlMock.NewRows([]string{"UserId", "Username", "Password"}).
			AddRow(mock.UserId, mock.Username, encryptedPassword),
	)

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  mock.Password,
		RequestId: uuid.New().String(),
		Token:     mock.Token,
	}
	userId, err := client.Login(mock.Ctx, request)
	if err != nil {
		t.Fatalf("error testing Login: %v", err)
	}
	assert.Equal(mock.UserId, userId.UserId)
}

func TestLogin_400UserDoesNotExist(t *testing.T) {
	assert := assert.New(t)
	client, sqlMock, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	sqlMock.ExpectQuery("SELECT * FROM `ACCOUNT` WHERE Username = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?").
		WithArgs(mock.Username, 1).WillReturnError(gorm.ErrRecordNotFound)

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  mock.Password,
		RequestId: uuid.New().String(),
		Token:     mock.Token,
	}
	userId, err := client.Login(mock.Ctx, request)
	assert.Equal(
		"rpc error: code = Code(400) desc = err: username or password is incorrect | fields: []",
		err.Error(),
	)
	assert.Nil(userId)
}

func TestLogin_400EncryptedKeyIsInvalid(t *testing.T) {
	assert := assert.New(t)
	client, sqlMock, closer := mock.StartAccountServer(mock.Ctx, t)
	defer closer()

	sqlMock.ExpectQuery(
		"SELECT * FROM `ACCOUNT` WHERE Username = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?",
	).WithArgs(mock.Username, 1).WillReturnRows(
		sqlMock.NewRows([]string{"UserId", "Username", "Password"}).
			AddRow(mock.UserId, mock.Username, "another_pasword"),
	)

	request := &pb.UserCredentialsRequest{
		Username:  mock.Username,
		Password:  mock.Password,
		RequestId: uuid.New().String(),
		Token:     mock.Token,
	}
	userId, err := client.Login(mock.Ctx, request)
	assert.Equal(
		"rpc error: code = Code(400) desc = err: username or password is incorrect | fields: []",
		err.Error(),
	)
	assert.Nil(userId)
}
