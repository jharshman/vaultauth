package main

import (
	"github.com/hashicorp/vault/api"
)

type Logical interface {
	Read(path string) (*api.Secret, error)
	Write(path string, data map[string]interface{}) (*api.Secret, error)
}

type Conn struct {
	client Logical
}

func (c *Conn) Auth(f func(c *Conn) (*api.Secret, error)) (map[string]interface{}, error) {
	s, err := f(c)
	if err != nil {
		return nil, err
	}
	return s.Data, nil
}
