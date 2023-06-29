package types

type (
	// Options represent the CLI options (flags and arguments) passed to the binary.
	Options struct {
		PortainerURL   *string
		Username       *string
		Password       *string
		ConfigFilePath *string
		InsecureTls    *bool
	}
)
