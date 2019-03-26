[![Go Report Card](https://goreportcard.com/badge/github.com/jharshman/vaultutil)](https://goreportcard.com/report/github.com/jharshman/vaultutil)
[![Build Status](https://travis-ci.org/jharshman/vaultutil.svg?branch=master)](https://travis-ci.org/jharshman/vaultutil)

# vaultutil

Simple auth utilities.

## Usage

Add the module:
`$ go get -m github.com/jharshman/vaultutil@v0.0.2`

Example:
```
// supply your own vault client to vaultutil.Conn{}
vClient, err := vault.NewClient(vault.DefaultConfig())
if err != nil {
	return err
}
vClient.SetAddress(h.vaultAddr)

conn := vaultutil.Conn{
	Client: vClient.Logical(),
}

// use vaultutil.Auth() to call one of the supported auth functions.
s, err := conn.Auth(vaultutil.Github(in.Token))
if err != nil {
	return err
}

vClient.SetToken(s.ClientToken)

// you can now use conn.Client to do whatever operations
// are permitted by the policy bound to the resulting ClientToken.
```

You may also write your own auth function.
```
cc.Auth(func(u, p string) func(c *Conn) (*api.Secret, error) {
	return func(c *Conn) (*api.Secret, error) {
		return &api.Secret{}, nil
	}
}("someuser", "somepassword"))
```
