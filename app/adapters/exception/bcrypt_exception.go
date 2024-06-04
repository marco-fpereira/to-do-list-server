package exception

import "fmt"

type BCryptException struct {
	StatusCode int
	Message    string
}

func (b *BCryptException) Error() string {
	return fmt.Sprintf("err: %s", b.Message)
}

func BuildBCryptException(message string) error {
	return &BCryptException{
		StatusCode: 500,
		Message:    message,
	}
}
