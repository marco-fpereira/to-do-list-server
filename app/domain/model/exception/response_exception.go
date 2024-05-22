package exception

import "fmt"

type ResponseException struct {
	StatusCode int
	Err        error
	Fields     []string
}

func (r *ResponseException) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
