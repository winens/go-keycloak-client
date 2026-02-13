package keycloak

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	config      *Config
	httpClient  *http.Client
	credentials *clientcredentials.Config
	adminApiUrl string
}

func NewClient(ctx context.Context, config *Config) *Client {
	credentials := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.BaseURL + "/realms/" + config.Realm + "/protocol/openid-connect/token",
	}

	httpClient := credentials.Client(ctx)

	return &Client{
		config:      config,
		credentials: &credentials,
		httpClient:  httpClient,
		adminApiUrl: config.BaseURL + "/admin/realms/" + config.Realm,
	}
}
