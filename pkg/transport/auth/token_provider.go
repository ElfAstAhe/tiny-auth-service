package auth

type TokenProvider interface {
	GetAccessToken() (string, error)
}
