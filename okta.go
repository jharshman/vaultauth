package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func Okta(u, p string) func(c *Conn) (*api.Secret, error) {
	return func(c *Conn) (*api.Secret, error) {
		return c.client.Write(fmt.Sprintf("/auth/okta/login/%s", u), map[string]interface{}{
			"password": p,
		})
	}
}
