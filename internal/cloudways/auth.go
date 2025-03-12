package cloudways

import (
	"fmt"
)

// AuthResponse defines the structure for the Cloudways OAuth token response.
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken authenticates the user with their email and API key,
// returning the access token if successful.
func GetAccessToken(email, apiKey string) (string, error) {
	// Initialize the HTTP client without any existing token.
	InitClient("")

	// Structure to store the token response
	var authResp AuthResponse

	// Send POST request to Cloudways OAuth endpoint
	resp, err := client.R().
		SetFormData(map[string]string{
			"email":   email,
			"api_key": apiKey,
		}).
		SetResult(&authResp). // Automatically unmarshal the JSON response into authResp
		Post("/oauth/access_token")

	// Handle request errors (network issues, etc.)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Cloudways API: %v", err)
	}

	// Handle invalid credentials or other API errors
	if resp.IsError() {
		return "", fmt.Errorf("error response: %v", resp.String())
	}

	// Log or display the token (for debug/testing purposes)
	fmt.Printf("✅ Access Token retrieved successfully: %s\n", authResp.AccessToken)

	// Return the access token
	return authResp.AccessToken, nil
}
