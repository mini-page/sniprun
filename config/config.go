package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GeminiAPIKey    string `yaml:"gemini_api_key"`
	AutoUpdate      bool   `yaml:"auto_update"`
	DefaultCategory string `yaml:"default_category"`
	SkipSecurity    bool   `yaml:"skip_security"`
}

var DefaultConfig = Config{
	GeminiAPIKey:    "",
	AutoUpdate:      false,
	DefaultCategory: "general",
	SkipSecurity:    false,
}

// Load reads config from ~/.sniprun/config.yaml
func Load(configDir string) (*Config, error) {
	configPath := filepath.Join(configDir, "config.yaml")

	// Return default if config doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &DefaultConfig, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Override with environment variables
	if apiKey := os.Getenv("GEMINI_API_KEY"); apiKey != "" {
		cfg.GeminiAPIKey = apiKey
	}

	return &cfg, nil
}

// Save writes config to ~/.sniprun/config.yaml
func Save(cfg *Config, configDir string) error {
	configPath := filepath.Join(configDir, "config.yaml")

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}