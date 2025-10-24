package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mini-page/sniprun/internal/security"
	"github.com/mini-page/sniprun/internal/snip"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [snip-name]",
	Short: "Create a new local snip",
	Long:  `Interactively create a new snip and save it to your local repository`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snipName := args[0]
		reader := bufio.NewReader(os.Stdin)

		_, _, err := snip.FindSnip(GetConfigDir(), snipName)
		if err == nil {
			fmt.Printf("Snip '%s' already exists. Overwrite? (yes/no): ", snipName)
			response, _ := reader.ReadString('\n')
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "yes" && response != "y" {
				fmt.Println("Cancelled")
				return
			}
		}

		fmt.Printf("Creating snip: %s\n\n", snipName)

		fmt.Print("Description: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		fmt.Print("Category (optional): ")
		category, _ := reader.ReadString('\n')
		category = strings.TrimSpace(category)

		fmt.Print("Command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		fmt.Print("Arguments (comma-separated, e.g., 'branch,message' or leave empty): ")
		argsInput, _ := reader.ReadString('\n')
		argsInput = strings.TrimSpace(argsInput)

		var argsList []string
		if argsInput != "" {
			for _, part := range strings.Split(argsInput, ",") {
				argsList = append(argsList, strings.TrimSpace(part))
			}
			fmt.Println("\nUse placeholders in command like: {{branch}}, {{message}}")
		}

		fmt.Println("\nValidating command security...")
		result, err := security.ValidateCommand(command)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Security check failed: %v\n", err)
		} else if result.RiskLevel == security.RiskDangerous {
			fmt.Fprintf(os.Stderr, "❌ Cannot add: This command appears dangerous\n")
			fmt.Fprintf(os.Stderr, "Reason: %s\n", result.Reason)
			os.Exit(1)
		} else if result.RiskLevel == security.RiskWarning {
			fmt.Printf("⚠️ Warning: %s\n", result.Reason)
			if !security.PromptUserConfirmation(command, "Add this snip anyway?") {
				fmt.Println("Cancelled")
				return
			}
		} else {
			fmt.Println("✓ Command validated")
		}

		s := &snip.Snip{
			Name:        snipName,
			Description: description,
			Command:     command,
			Args:        argsList,
			Category:    category,
			Trust:       "local",
		}

		localPath := filepath.Join(GetConfigDir(), "snips", "local", snipName+".yaml")
		if err := snip.SaveSnip(s, localPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving snip: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n✓ Snip '%s' created successfully\n", snipName)
		fmt.Printf("Run: sniprun %s", snipName)
		if len(argsList) > 0 {
			fmt.Printf(" <%s>", strings.Join(argsList, "> <"))
		}
		fmt.Println()
	},
}
