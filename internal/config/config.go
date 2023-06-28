package config

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
