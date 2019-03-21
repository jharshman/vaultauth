package vaultutil

import (
	"github.com/hashicorp/vault/api"
)

// Github returns a function facilitating the Github auth flow.
// Takes a GitHub personal access token as a parameter.
// Ref: https://www.vaultproject.io/docs/auth/github.html
func Github(t string) func(c *Conn) (*api.Secret, error) {
	return func(c *Conn) (*api.Secret, error) {
		return c.Client.Write("/auth/github/login", map[string]interface{}{
			"token": t,
		})
	}
}
