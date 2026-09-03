package agents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Struct untuk penampung JSON response dari API
type MakersResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Struct Agent / Service Client
type MakersAgent struct {
	ApiKey  string
	BaseURL string
	Model   string
}

// Constructor untuk membuat instance agent baru
func NewMakersAgent() *MakersAgent {
	return &MakersAgent{
		ApiKey:  os.Getenv("API_KEY"),
		BaseURL: strings.TrimRight(os.Getenv("BASE_URL"), "/"),
		Model:   os.Getenv("MAKERS_MODEL"),
	}
}

// Method untuk berkomunikasi dengan EdgeOne API
func (a *MakersAgent) TestMakers(systemPrompt, userPrompt string) (string, error) {
	// Pastikan endpoint URL terbentuk dengan benar
	targetURL := fmt.Sprintf("%s/chat/completions", a.BaseURL)

	requestBody := map[string]interface{}{
		"model": a.Model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt,
			},
			{
				"role":    "user",
				"content": userPrompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+a.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Timeout ditambahkan untuk mencegah request menggantung terlalu lama
	client := &http.Client{Timeout: 60 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("makers API error: status=%d body=%s", resp.StatusCode, string(body))
	}

	var result MakersResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("AI tidak memberikan response")
	}

	return result.Choices[0].Message.Content, nil
}