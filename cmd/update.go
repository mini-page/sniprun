package cmd

import (
	"fmt"
	"os"

	"github.com/mini-page/sniprun/internal/repo"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Sync community snips from GitHub",
	Long:  `Download and update community-contributed snips from the official repository`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating community snips...")

		communityDir := GetConfigDir() + "/snips/community"

		count, err := repo.SyncCommunitySnips(communityDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error syncing snips: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✓ Successfully synced %d community snips\n", count)
		fmt.Println("Run 'sniprun list' to see available snips")
	},
}
