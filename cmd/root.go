package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	configDir string
	rootCmd   = &cobra.Command{
		Use:   "sniprun",
		Short: "Run complex commands with short, memorable snips",
		Long:  `sniprun - Execute simplified aliases for complex commands with community contributions`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	
	// Global flags
	rootCmd.PersistentFlags().StringVar(&configDir, "config", "", "config directory (default is $HOME/.sniprun)")
}

func initConfig() {
	if configDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
			os.Exit(1)
		}
		configDir = filepath.Join(home, ".sniprun")
	}

	// Create config directories if they don't exist
	dirs := []string{
		configDir,
		filepath.Join(configDir, "snips", "local"),
		filepath.Join(configDir, "snips", "community"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}
}

func GetConfigDir() string {
	return configDir
}