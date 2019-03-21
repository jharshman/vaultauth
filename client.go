package vaultutil

import (
	"github.com/hashicorp/vault/api"
)

// Vault api.Logical implements this interface.
// Useful for generating mocks.
type Logical interface {
	Write(path string, data map[string]interface{}) (*api.Secret, error)
}

type Conn struct {
	Client Logical
}

// Auth wraps different authentication functions. It executes the function pointed to by f,
// which is in turn a function facilitating some sort of authentication.
// It returns a vault api.SecretAuth structure and an error.
func (c *Conn) Auth(f func(c *Conn) (*api.Secret, error)) (*api.SecretAuth, error) {
	s, err := f(c)
	if err != nil {
		return nil, err
	}
	return s.Auth, nil
}
