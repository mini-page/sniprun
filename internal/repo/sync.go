package repo

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// GitHub repository for community snips
	RepoOwner = "sniprun"
	RepoName  = "snips"
	RepoBranch = "main"
)

// SyncCommunitySnips downloads community snips from GitHub
func SyncCommunitySnips(targetDir string) (int, error) {
	// Download zip archive from GitHub
	zipURL := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", 
		RepoOwner, RepoName, RepoBranch)

	fmt.Printf("Downloading from: %s\n", zipURL)

	resp, err := http.Get(zipURL)
	if err != nil {
		return 0, fmt.Errorf("failed to download repository: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to download: HTTP %d", resp.StatusCode)
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "sniprun-*.zip")
	if err != nil {
		return 0, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Download to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return 0, fmt.Errorf("failed to save download: %w", err)
	}

	// Extract YAML files
	count, err := extractSnips(tmpFile.Name(), targetDir)
	if err != nil {
		return 0, fmt.Errorf("failed to extract snips: %w", err)
	}

	return count, nil
}

// extractSnips extracts .yaml files from zip to target directory
func extractSnips(zipPath, targetDir string) (int, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	count := 0

	// Clean target directory first
	if err := os.RemoveAll(targetDir); err != nil {
		return 0, err
	}
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return 0, err
	}

	for _, f := range r.File {
		// Only extract .yaml files
		if !strings.HasSuffix(f.Name, ".yaml") {
			continue
		}

		// Skip directories and hidden files
		if f.FileInfo().IsDir() || strings.Contains(filepath.Base(f.Name), ".") && !strings.HasSuffix(f.Name, ".yaml") {
			continue
		}

		// Get just the filename
		filename := filepath.Base(f.Name)
		targetPath := filepath.Join(targetDir, filename)

		// Extract file
		if err := extractFile(f, targetPath); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to extract %s: %v\n", filename, err)
			continue
		}

		count++
	}

	return count, nil
}

func extractFile(f *zip.File, targetPath string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	outFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	return err
}