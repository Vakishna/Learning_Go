// auth/auth.go

package auth

import (
	"context"
	"log"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://ttech.au.auth0.com/")
	if err != nil {
		log.Printf("Failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "gZk605RjXRmxmrjGP7DkO5WbNMJENKDH",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "http://localhost:3000/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil

}
