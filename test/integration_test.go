package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestSniprunCLI(t *testing.T) {
	cmd := exec.Command("go", "run", "../../main.go", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("sniprun command failed: %v\nOutput: %s", err, output)
	}

	if !strings.Contains(string(output), "Available Snips") {
		t.Errorf("Expected 'Available Snips' in output, got: %s", output)
	}
}
