package dto

// ChangedKeysDTO encapsulates the outbound payload carrying structural asymmetric
// cryptographic key tokens returned upon a successful key rotation event stream execution.
type ChangedKeysDTO struct {
	PublicKey  string `json:"public_key,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
} // @name ChangedKeysDTO
