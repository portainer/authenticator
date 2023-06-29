package config

import (
	kingpin "github.com/alecthomas/kingpin/v2"
	"github.com/portainer/authenticator/internal/types"
)

// ConfigFilePath represent the path to the config.json file that will be updated.
const DefaultConfigFilePath = "/config.json"

// ParseOptions parses the arguments/flags passed to the binary.
func ParseOptions() *types.Options {
	options := types.Options{
		ConfigFilePath: kingpin.Flag("config", "Path to the configuration file to update").Default(DefaultConfigFilePath).Short('c').String(),
		PortainerURL:   kingpin.Arg("url", "URL of the Portainer instance.").Required().String(),
		Username:       kingpin.Arg("username", "Portainer Username").Required().String(),
		Password:       kingpin.Arg("password", "Portainer Password").Required().String(),
		InsecureTls:    kingpin.Arg("insecureTls", "Disable TLS certificate verification. Defaults to false").Default("false").Bool(),
	}

	kingpin.Parse()

	return &options
}
