package output

type AuthenticationPort interface {
	ValidateToken(token string) (bool, error)
	ValidateClaim(
		token string,
		claimName string,
		claimValue string,
	) (bool, error)
}
