package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/portainer/authenticator/internal/api"
	"github.com/portainer/authenticator/internal/config"
)

func main() {
	// parse options from the config file
	options := config.ParseOptions()

	// initialise a HTTP client to communicate with the Portainer API
	client, err := api.NewPortainerApiClientModel(*options.PortainerURL, *options.Username, *options.Password, *options.InsecureTls)
	if err != nil {
		log.Fatalf("unable to create portainer client: %s...", err.Error())
	}

	jwt, err := api.GetAuthJwt(client, options)
	if err != nil {
		log.Fatalf("unable to retrieve portainer jwt: %s... please check your config file...", err.Error())
	}

	raw, err := os.ReadFile(*options.ConfigFilePath)
	if err != nil {
		log.Fatalf("Unable to read configuration file: %s", err.Error())
	}

	var fileData map[string]interface{}
	err = json.Unmarshal(raw, &fileData)
	if err != nil {
		log.Fatalf("Unable to decode configuration file: %s", err.Error())
	}

	if fileData["HttpHeaders"] == nil {
		fileData["HttpHeaders"] = make(map[string]interface{})
	}

	headersObject := fileData["HttpHeaders"].(map[string]interface{})
	authorizationHeaderValue := "Bearer " + jwt
	headersObject["Authorization"] = authorizationHeaderValue

	buf, err := json.MarshalIndent(fileData, "", "  ")
	if err != nil {
		log.Fatalf("Unable to encode configuration file content: %s", err.Error())
	}

	err = os.WriteFile(*options.ConfigFilePath, buf, 0644)
	if err != nil {
		log.Fatalf("Unable to write to configuration file: %s", err.Error())
	}
}
