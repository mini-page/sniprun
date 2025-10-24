package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/mini-page/sniprun/internal/snip"

	"github.com/spf13/cobra"
)

var (
	listCategory string
)

func init() {
	listCmd.Flags().StringVarP(&listCategory, "category", "c", "", "Filter by category")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available snips",
	Long:  `Display all installed snips from local and community repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		snips, err := snip.ListSnips(GetConfigDir())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading snips: %v\n", err)
			os.Exit(1)
		}

		if len(snips) == 0 {
			fmt.Println("No snips found. Run 'sniprun update' to fetch community snips or 'sniprun add' to create your own.")
			return
		}

		// Sort by name
		names := make([]string, 0, len(snips))
		for name := range snips {
			names = append(names, name)
		}
		sort.Strings(names)

		// Group by category
		categories := make(map[string][]string)
		for _, name := range names {
			s := snips[name]
			
			// Filter by category if specified
			if listCategory != "" && s.Category != listCategory {
				continue
			}

			cat := s.Category
			if cat == "" {
				cat = "uncategorized"
			}
			categories[cat] = append(categories[cat], name)
		}

		// Display
		fmt.Printf("Available snips (%d total):\n\n", len(snips))

		catNames := make([]string, 0, len(categories))
		for cat := range categories {
			catNames = append(catNames, cat)
		}
		sort.Strings(catNames)

		for _, cat := range catNames {
			fmt.Printf("── %s ──\n", strings.ToUpper(cat))
			for _, name := range categories[cat] {
				s := snips[name]
				trustIcon := ""
				switch s.Trust {
				case "local":
					trustIcon = "🔧"
				case "community":
					trustIcon = "🌐"
				case "verified":
					trustIcon = "✓"
				}

				argsStr := ""
				if len(s.Args) > 0 {
					argsStr = fmt.Sprintf(" [%s]", strings.Join(s.Args, ", "))
				}

				fmt.Printf("  %s %s%s\n", trustIcon, name, argsStr)
				fmt.Printf("     %s\n", s.Description)
			}
			fmt.Println()
		}

		fmt.Println("Run 'sniprun explain <name>' to see the command")
		fmt.Println("Run 'sniprun <name> [args]' to execute")
	},
}