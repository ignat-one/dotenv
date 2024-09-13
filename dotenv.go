package dotenv

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load locates and loads the .env file based on the provided options.
func Load(options ...Option) error {
	config := NewConfig(options)
	envFile, err := findEnvFile(config)
	if err != nil {
		return fmt.Errorf("failed to retrieve environment file: %w", err)
	}
	err = godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("failed to load environment file: %w", err)
	}
	return nil
}

// findEnvFile returns the full path of the .env file it finds based on the anchor file.
func findEnvFile(config *Config) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	for {
		anchorFilePath := filepath.Join(currentDir, config.AnchorFile)
		if _, err := os.Stat(anchorFilePath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			return "", fmt.Errorf("anchor file %s not found", config.AnchorFile)
		}

		currentDir = parent
	}

	return filepath.Join(currentDir, config.EnvFilename), nil
}
