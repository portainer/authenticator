package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/portainer/authenticator/internal/cli"
)

type authenticationRequestPayload struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func main() {
	options := cli.ParseOptions()

	apiURL, err := url.Parse(*options.PortainerAPI)
	if err != nil {
		log.Fatalf("Invalid Portainer URL: %s", err.Error())
	}

	apiURL.Path = path.Join(apiURL.Path, "/api/auth")
	authenticationURL := apiURL.String()

	payload := authenticationRequestPayload{
		Username: *options.Username,
		Password: *options.Password,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Unable to encode payload: %s", err.Error())
	}

	response, err := http.Post(authenticationURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Unable to execute authentication request: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusUnprocessableEntity {
			log.Fatalf("Invalid credentials: %s / %s", payload.Username, payload.Password)
		} else {
			log.Fatalf("An error occured during HTTP authentication")
		}
	}

	data, err := getResponseBodyAsJSONObject(response)
	if err != nil {
		log.Fatalf("Unable to read authentication response: %s", err.Error())
	}

	token := data["jwt"]

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
	authorizationHeaderValue := "Bearer " + token.(string)
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

func getResponseBodyAsJSONObject(response *http.Response) (map[string]interface{}, error) {
	var data map[string]interface{}

	err := json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
