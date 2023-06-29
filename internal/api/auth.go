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

// func GetAuthType(portainerHostname string, portainerUsername string, portainerPassword string, portainerApiKey string, insecureTls bool) (runtime.ClientAuthInfoWriter, error) {
// 	if (portainerApiKey != "") || (portainerUsername != "" && portainerPassword != "" && portainerApiKey != "") {
// 		return httptransport.APIKeyAuth("X-API-KEY", "header", portainerApiKey), nil
// 	} else if portainerUsername != "" && portainerPassword != "" && portainerApiKey == "" {
// 		var token, err = GetAuthJwt(portainerHostname, portainerUsername, portainerPassword, insecureTls)
// 		return httptransport.BearerToken(token), err
// 	} else {
// 		return nil, errors.New("no authentication type was found. Ensure it is either basic using username and password, or API Key based")
// 	}
// }
