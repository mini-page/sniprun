package security

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type RiskLevel string

const (
	RiskSafe      RiskLevel = "safe"
	RiskWarning   RiskLevel = "warning"
	RiskDangerous RiskLevel = "dangerous"
)

type ValidationResult struct {
	Safe      bool
	RiskLevel RiskLevel
	Reason    string
}

// ValidateCommand uses Gemini API to check if a command is potentially harmful
func ValidateCommand(command string) (*ValidationResult, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		// Skip validation if no API key is set
		return &ValidationResult{
			Safe:      true,
			RiskLevel: RiskSafe,
			Reason:    "Security validation skipped (no API key)",
		}, nil
	}

	// Prepare Gemini API request
	prompt := fmt.Sprintf(`Analyze this shell command for security risks:
Command: %s

Respond ONLY with JSON in this format:
{
  "risk_level": "safe|warning|dangerous",
  "reason": "brief explanation"
}

Risk levels:
- safe: Normal operation, no risk
- warning: Could be destructive (rm, format, etc.) but legitimate
- dangerous: Malicious intent detected (exfiltration, malware, etc.)`, command)

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": prompt},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"temperature": 0.1,
			"maxOutputTokens": 200,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Call Gemini API
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s", apiKey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var apiResp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(apiResp.Candidates) == 0 || len(apiResp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("empty response from API")
	}

	// Extract JSON from response
	responseText := apiResp.Candidates[0].Content.Parts[0].Text
	responseText = strings.TrimPrefix(responseText, "```json")
	responseText = strings.TrimSuffix(responseText, "```")
	responseText = strings.TrimSpace(responseText)

	var result struct {
		RiskLevel string `json:"risk_level"`
		Reason    string `json:"reason"`
	}

	if err := json.Unmarshal([]byte(responseText), &result); err != nil {
		return nil, fmt.Errorf("failed to parse validation result: %w", err)
	}

	// Convert to ValidationResult
	vr := &ValidationResult{
		Safe:      result.RiskLevel == "safe",
		RiskLevel: RiskLevel(result.RiskLevel),
		Reason:    result.Reason,
	}

	return vr, nil
}

// PromptUserConfirmation asks user to confirm execution of risky commands
func PromptUserConfirmation(command string, reason string) bool {
	fmt.Printf("\n⚠️  WARNING: This command may be risky\n")
	fmt.Printf("Command: %s\n", command)
	fmt.Printf("Reason: %s\n\n", reason)
	fmt.Print("Do you want to continue? (yes/no): ")

	var response string
	fmt.Scanln(&response)

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "yes" || response == "y"
}