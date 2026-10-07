package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrSessionExpired     = errors.New("session expired")
	ErrSessionRevoked     = errors.New("session revoked")
)

type User struct {
	ID        string
	Email     string
	Phone     string
	Role      string
	FirstName string
	LastName  string
}

type Session struct {
	Token     string
	ExpiresAt string
	User      User
}
