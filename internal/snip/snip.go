package snip

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Snip struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Command     string   `yaml:"command"`
	Args        []string `yaml:"args"`
	Category    string   `yaml:"category"`
	Trust       string   `yaml:"trust"` // community | local | verified
}

// LoadSnip reads a snip from a YAML file
func LoadSnip(path string) (*Snip, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read snip: %w", err)
	}

	var snip Snip
	if err := yaml.Unmarshal(data, &snip); err != nil {
		return nil, fmt.Errorf("failed to parse snip: %w", err)
	}

	return &snip, nil
}

// SaveSnip writes a snip to a YAML file
func SaveSnip(snip *Snip, path string) error {
	data, err := yaml.Marshal(snip)
	if err != nil {
		return fmt.Errorf("failed to marshal snip: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write snip: %w", err)
	}

	return nil
}

// InterpolateArgs replaces {{arg}} placeholders with actual values
func (s *Snip) InterpolateArgs(args []string) (string, error) {
	command := s.Command

	// Check if we have the right number of arguments
	if len(args) != len(s.Args) {
		return "", fmt.Errorf("expected %d arguments (%v), got %d", len(s.Args), s.Args, len(args))
	}

	// Replace placeholders
	for i, argName := range s.Args {
		placeholder := fmt.Sprintf("{{%s}}", argName)
		command = strings.ReplaceAll(command, placeholder, args[i])
	}

	return command, nil
}

// ListSnips returns all available snips from local and community directories
func ListSnips(configDir string) (map[string]*Snip, error) {
	snips := make(map[string]*Snip)

	// Load from both directories
	dirs := []string{
		filepath.Join(configDir, "snips", "local"),
		filepath.Join(configDir, "snips", "community"),
	}

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue // Skip if directory doesn't exist
		}

		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".yaml") {
				continue
			}

			path := filepath.Join(dir, entry.Name())
			snip, err := LoadSnip(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: failed to load %s: %v\n", path, err)
				continue
			}

			// Use snip name as key, local overrides community
			snips[snip.Name] = snip
		}
	}

	return snips, nil
}

// FindSnip locates a snip by name
func FindSnip(configDir, name string) (*Snip, string, error) {
	// Check local first
	localPath := filepath.Join(configDir, "snips", "local", name+".yaml")
	if _, err := os.Stat(localPath); err == nil {
		snip, err := LoadSnip(localPath)
		return snip, localPath, err
	}

	// Check community
	communityPath := filepath.Join(configDir, "snips", "community", name+".yaml")
	if _, err := os.Stat(communityPath); err == nil {
		snip, err := LoadSnip(communityPath)
		return snip, communityPath, err
	}

	return nil, "", fmt.Errorf("snip '%s' not found", name)
}