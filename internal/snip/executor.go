package snip

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Execute runs the snip command in a subprocess
func (s *Snip) Execute(args []string, dryRun bool) error {
	command, err := s.InterpolateArgs(args)
	if err != nil {
		return err
	}

	if dryRun {
		fmt.Printf("Would execute: %s\n", command)
		return nil
	}

	fmt.Printf("Executing: %s\n", command)

	// Determine shell based on OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	// Connect to stdio
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// ExecuteInShell runs the command and returns output for evaluation in current shell
// This is for advanced use cases like cd, export, etc.
func (s *Snip) ExecuteInShell(args []string) (string, error) {
	command, err := s.InterpolateArgs(args)
	if err != nil {
		return "", err
	}

	fmt.Fprintf(os.Stderr, "Note: Use 'eval $(sniprun %s --source)' to execute in current shell\n", s.Name)
	return command, nil
}