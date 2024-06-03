package output

type AuthenticationPort interface {
	ValidateToken(token string) (bool, error)
}
