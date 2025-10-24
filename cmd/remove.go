package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sniprun/internal/security"
	"sniprun/internal/snip"

	"github.com/spf13/cobra"
)

var forceRemove bool

func init() {
	removeCmd.Flags().BoolVarP(&forceRemove, "force", "f", false, "Skip confirmation prompt")
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [snip-name]",
	Short: "Delete a local snip",
	Long:  `Remove a snip from your local repository. Community snips cannot be removed.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snipName := args[0]

		// Find snip
		s, path, err := snip.FindSnip(GetConfigDir(), snipName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Only allow removing local snips
		if s.Trust != "local" {
			fmt.Fprintf(os.Stderr, "Error: Cannot remove %s snips. Only local snips can be removed.\n", s.Trust)
			fmt.Fprintf(os.Stderr, "To hide community snips, delete them from: %s\n", filepath.Join(GetConfigDir(), "snips", "community"))
			os.Exit(1)
		}

		// Security check for deletion
		if !forceRemove {
			fmt.Printf("About to remove snip: %s\n", s.Name)
			fmt.Printf("Description: %s\n", s.Description)
			fmt.Printf("Command: %s\n\n", s.Command)

			// Validate deletion
			result, err := security.ValidateCommand(fmt.Sprintf("rm %s", path))
			if err == nil && result.RiskLevel == security.RiskDangerous {
				fmt.Fprintf(os.Stderr, "Security warning: %s\n", result.Reason)
			}

			if !security.PromptUserConfirmation(s.Command, "Remove this snip?") {
				fmt.Println("Cancelled")
				return
			}
		}

		// Delete file
		if err := os.Remove(path); err != nil {
			fmt.Fprintf(os.Stderr, "Error removing snip: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✓ Snip '%s' removed successfully\n", snipName)
	},
}