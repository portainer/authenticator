package cli

import (
	kingpin "github.com/alecthomas/kingpin/v2"
	"github.com/portainer/authenticator/internal/config"
)

// ParseOptions parses the arguments/flags passed to the binary.
func ParseOptions() *config.Options {

	options := &config.Options{
		ConfigFilePath: kingpin.Flag("config", "Path to the configuration file to update").Default(config.DefaultConfigFilePath).Short('c').String(),
		PortainerAPI:   kingpin.Arg("portainer API URL", "URL of the Portainer API.").Required().String(),
		Username:       kingpin.Arg("Username", "Username").Required().String(),
		Password:       kingpin.Arg("Password", "Password").Required().String(),
	}

	kingpin.Parse()

	return options
}
