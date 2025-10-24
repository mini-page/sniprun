package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mini-page/sniprun/internal/snip"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(explainCmd)
}

var explainCmd = &cobra.Command{
	Use:   "explain [snip-name]",
	Short: "Show what a snip will execute",
	Long:  `Display detailed information about a snip without executing it`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snipName := args[0]

		s, path, err := snip.FindSnip(GetConfigDir(), snipName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Snip: %s\n", s.Name)
		fmt.Printf("Description: %s\n", s.Description)
		fmt.Printf("Category: %s\n", s.Category)
		fmt.Printf("Trust: %s\n", s.Trust)
		fmt.Printf("Path: %s\n\n", path)

		fmt.Println("Command:")
		fmt.Printf("  %s\n\n", s.Command)

		if len(s.Args) > 0 {
			fmt.Printf("Arguments required: %s\n", strings.Join(s.Args, ", "))
			fmt.Printf("Usage: sniprun %s <%s>\n", s.Name, strings.Join(s.Args, "> <"))

			fmt.Println("\nExample with placeholders:")
			exampleArgs := make([]string, len(s.Args))
			for i, arg := range s.Args {
				exampleArgs[i] = fmt.Sprintf("<your-%s>", arg)
			}
			command, _ := s.InterpolateArgs(exampleArgs)
			fmt.Printf("  %s\n", command)
		} else {
			fmt.Printf("Usage: sniprun %s\n", s.Name)
		}
	},
}
