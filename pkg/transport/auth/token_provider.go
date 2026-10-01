package auth

// TokenProvider defines the strongly-typed contract for securely acquiring authorization identifiers.
// It abstracts the application delivery and transport layers from concrete security token generation mechanics,
// such as ephemeral in-memory JWT assembly or remote upstream OAuth2/OIDC provider exchange sequences.
type TokenProvider interface {
	// GetAccessToken yields a valid cryptographically signed access token string.
	// Returns a structured error state if the provider fails to generate, refresh, or fetch the token payload.
	GetAccessToken() (string, error)
}
