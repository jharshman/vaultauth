package vaultutil

import (
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/vault/api"
	"github.com/jharshman/vaultauth/mock"
	"github.com/matryer/is"
	"testing"
)

var writeResult = &api.Secret{
	Auth: &api.SecretAuth{
		ClientToken: "c4f280f6-fdb2-18eb-89d3-589e2e834cdb",
		Policies:    []string{"policy1", "policy2"},
		Metadata: map[string]string{
			"foo": "bar",
			"fiz": "buz",
		},
	},
}

func TestConn_OktaAuth(t *testing.T) {

	ctrl := gomock.NewController(t)
	vcli := mock.NewMockLogical(ctrl)

	vcli.EXPECT().Write(gomock.Any(), gomock.Any()).Return(writeResult, nil)

	cc := &Conn{
		client: vcli,
	}

	is := is.New(t)

	result, err := cc.Auth(Okta("username", "password"))
	is.NoErr(err)
	is.Equal(result, writeResult.Auth)
}

func TestConn_GithubAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	vcli := mock.NewMockLogical(ctrl)

	vcli.EXPECT().Write("/auth/github/login", map[string]interface{}{
		"token": "sometoken",
	}).Return(writeResult, nil)

	cc := &Conn{
		client: vcli,
	}

	is := is.New(t)

	result, err := cc.Auth(Github("sometoken"))
	is.NoErr(err)
	is.Equal(result, writeResult.Auth)
}

func TestConn_JWTAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	vcli := mock.NewMockLogical(ctrl)

	vcli.EXPECT().Write("/auth/jwt/login", map[string]interface{}{
		"jwt":  "sometoken",
		"role": "somerole",
	}).Return(writeResult, nil)

	cc := &Conn{
		client: vcli,
	}

	is := is.New(t)

	result, err := cc.Auth(JWT("sometoken", "somerole"))
	is.NoErr(err)
	is.Equal(result, writeResult.Auth)
}
