package cloudways

import "github.com/go-resty/resty/v2"

const baseURL = "https://api.cloudways.com/api/v1"

var client *resty.Client

func InitClient(authToken string) {
	client = resty.New()
	client.SetBaseURL(baseURL)

	if authToken != "" {
		client.SetAuthToken(authToken)
	}
}
