package config

import (
	kingpin "github.com/alecthomas/kingpin/v2"
)

type (
	// Options represent the CLI options (flags and arguments) passed to the binary.
	Options struct {
		PortainerAPI   *string
		Username       *string
		Password       *string
		ConfigFilePath *string
	}
)

// ConfigFilePath represent the path to the config.json file that will be updated.
const DefaultConfigFilePath = "/config.json"

// ParseOptions parses the arguments/flags passed to the binary.
func ParseOptions() *Options {

	options := Options{
		ConfigFilePath: kingpin.Flag("config", "Path to the configuration file to update").Default(DefaultConfigFilePath).Short('c').String(),
		PortainerAPI:   kingpin.Arg("portainer API URL", "URL of the Portainer API.").Required().String(),
		Username:       kingpin.Arg("Username", "Username").Required().String(),
		Password:       kingpin.Arg("Password", "Password").Required().String(),
	}

	kingpin.Parse()

	return &options
}
