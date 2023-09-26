package api

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/portainer/authenticator/internal/types"
	"github.com/portainer/client-api-go/v2/client/auth"
	"github.com/portainer/client-api-go/v2/models"
)

func GetAuthJwt(r *PortainerApiClientModel, options *types.Options) (string, error) {
	response, err := r.Client.Auth.AuthenticateUser(&auth.AuthenticateUserParams{
		Context:    context.Background(),
		HTTPClient: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: *options.InsecureTls}}},
		Body: &models.AuthAuthenticatePayload{
			Username: options.Username,
			Password: options.Password,
		},
	})

	if err != nil {
		return "", err
	}

	return response.Payload.Jwt, err
}
