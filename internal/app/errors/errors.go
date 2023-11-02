package errors

import "fmt"

var (
	NotFound             = fmt.Errorf("not found")
	WrongEmailPassword   = fmt.Errorf("wrong email or password")
	EmptyAuthHeader      = fmt.Errorf("empty auth header")
	InvalidAuthHeader    = fmt.Errorf("invalid auth header")
	InvalidSigningMethod = fmt.Errorf("invalid signing method")
	TokenClaimsErr       = fmt.Errorf("token claims are not of type *tokenClaims")
)
