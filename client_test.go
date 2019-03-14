package main

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/vault/api"
	"github.com/jharshman/vaultauth/mock"
	"github.com/matryer/is"
	"testing"
)

var oktaResult = &api.Secret{
	Data: map[string]interface{}{
		"foo": "bar",
	},
}

func TestConn_OktaAuth(t *testing.T) {

	ctrl := gomock.NewController(t)
	vcli := mock.NewMockLogical(ctrl)

	vcli.EXPECT().Write(gomock.Any(), gomock.Any()).Return(oktaResult, nil)

	cc := &Conn{
		client: vcli,
	}

	is := is.New(t)

	result, err := cc.Auth(Okta("username", "password"))
	fmt.Printf("%v\n", result)
	is.NoErr(err)
}
