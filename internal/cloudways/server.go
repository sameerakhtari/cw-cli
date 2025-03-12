package cloudways

import (
	"fmt"
)

type Server struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	PublicIP string `json:"public_ip"`
	Status   string `json:"status"`
	Cloud    string `json:"cloud"`
	Region   string `json:"region"`
}

type GetServersResponse struct {
	Servers []Server `json:"servers"`
}

func ListServers() ([]Server, error) {
	// Load the config to get the AccessToken
	cfg, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	// Initialize the client with the saved token
	InitClient(cfg.AccessToken)

	// Call the Cloudways API
	var serversResp GetServersResponse

	resp, err := client.R().
		SetResult(&serversResp).
		Get("/server")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error fetching servers: %v", resp.String())
	}

	return serversResp.Servers, nil
}
