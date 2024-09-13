package dotenv

const (
	defaultAnchorFile  = "go.mod"
	defaultEnvFilename = ".env"
)

// Config holds configuration for locating the .env file.
type Config struct {
	AnchorFile  string
	EnvFilename string
}

// Option represents an option for configuring the dotenv loader.
type Option func(*Config)

// NewConfig creates a new dotenv configuration with provided options.
func NewConfig(options []Option) *Config {
	conf := &Config{
		AnchorFile:  defaultAnchorFile,
		EnvFilename: defaultEnvFilename,
	}
	for _, option := range options {
		option(conf)
	}
	return conf
}

// WithAnchorFile specifies a custom anchor file.
func WithAnchorFile(filename string) Option {
	return func(c *Config) {
		c.AnchorFile = filename
	}
}

// WithEnvFilename specifies a custom .env filename.
func WithEnvFilename(filename string) Option {
	return func(c *Config) {
		c.EnvFilename = filename
	}
}
