package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

// Okta returns a function facilitating the Okta auth flow.
// Takes a username and password as parameters.
// Ref: https://www.vaultproject.io/docs/auth/okta.html
func Okta(u, p string) func(c *Conn) (*api.Secret, error) {
	return func(c *Conn) (*api.Secret, error) {
		return c.client.Write(fmt.Sprintf("/auth/okta/login/%s", u), map[string]interface{}{
			"password": p,
		})
	}
}
