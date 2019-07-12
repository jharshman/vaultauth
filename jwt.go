package vaultauth

import "github.com/hashicorp/vault/api"

// JWT returns a function facilitating the JWT auth flow.
// Takes a JWT token and role as parameters.
// Ref: https://www.vaultproject.io/docs/auth/jwt.html
func JWT(t, r string) func(c *Conn) (*api.Secret, error) {
	return func(c *Conn) (*api.Secret, error) {
		return c.Client.Write("/auth/jwt/login", map[string]interface{}{
			"jwt":  t,
			"role": r,
		})
	}
}
