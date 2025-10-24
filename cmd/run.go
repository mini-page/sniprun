package cmd

import (
	"fmt"
	"os"

	"github.com/mini-page/sniprun/internal/security"
	"github.com/mini-page/sniprun/internal/snip"

	"github.com/spf13/cobra"
)

var (
	sourceMode bool
	skipSecurityCheck bool
)

func init() {
	runCmd.Flags().BoolVar(&sourceMode, "source", false, "Output command for shell evaluation (use with eval)")
	runCmd.Flags().BoolVar(&skipSecurityCheck, "skip-check", false, "Skip security validation")
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run [snip-name] [args...]",
	Short: "Execute a snip",
	Long:  `Execute a stored command snip with optional arguments`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snipName := args[0]
		snipArgs := args[1:]

		// Find the snip
		s, _, err := snip.FindSnip(GetConfigDir(), snipName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Run 'sniprun list' to see available snips\n")
			os.Exit(1)
		}

		// Interpolate arguments
		command, err := s.InterpolateArgs(snipArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Security validation (unless skipped)
		if !skipSecurityCheck {
			result, err := security.ValidateCommand(command)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Security check failed: %v\n", err)
				fmt.Fprintf(os.Stderr, "Continuing anyway (use --skip-check to suppress this warning)\n")
			} else if result.RiskLevel == security.RiskDangerous {
				fmt.Fprintf(os.Stderr, "❌ BLOCKED: This command appears dangerous\n")
				fmt.Fprintf(os.Stderr, "Reason: %s\n", result.Reason)
				fmt.Fprintf(os.Stderr, "Command: %s\n", command)
				os.Exit(1)
			} else if result.RiskLevel == security.RiskWarning {
				if !security.PromptUserConfirmation(command, result.Reason) {
					fmt.Println("Execution cancelled")
					os.Exit(0)
				}
			}
		}

		// Execute
		if sourceMode {
			fmt.Println(command)
		} else {
			if err := s.Execute(snipArgs, false); err != nil {
				fmt.Fprintf(os.Stderr, "Execution failed: %v\n", err)
				os.Exit(1)
			}
		}
	},
}