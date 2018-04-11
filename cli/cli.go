package cli

import (
	"github.com/portainer/authenticator"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// ParseOptions parses the arguments/flags passed to the binary.
func ParseOptions() *authenticator.Options {

	options := &authenticator.Options{
		ConfigFilePath: kingpin.Flag("config", "Path to the configuration file to update").Default(authenticator.DefaultConfigFilePath).Short('c').String(),
		PortainerAPI:   kingpin.Arg("portainer API URL", "URL of the Portainer API.").Required().String(),
		Username:       kingpin.Arg("Username", "Username").Required().String(),
		Password:       kingpin.Arg("Password", "Password").Required().String(),
	}

	kingpin.Parse()

	return options
}
