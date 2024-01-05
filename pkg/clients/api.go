package clients

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	baseURLEnv           = "API_BASEURL"
	usernameEnv          = "API_USERNAME"
	passwordEnv          = "API_PASSWORD"
	proxyHTTP            = "API_HTTP_PROXY"
	authenticationHeader = "Authorization"
)

type APIClient struct {
	client *resty.Client
}

type Settings struct {
	BaseURL  string `yaml:"baseurl"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	ProxyUrl string `yaml:"proxy"`
}

func NewSettings(baseURL, username, password, proxyURL string) (*Settings, error) {
	settings := &Settings{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		ProxyUrl: proxyURL,
	}

	// Do we have username, password and baseURL?
	if settings.Username == "" {
		return nil, ErrNoUsernameProvided
	}
	if settings.Password == "" {
		return nil, ErrNoPasswordProvided
	}
	if settings.BaseURL == "" {
		return nil, ErrNoBaseUrlProvided
	}

	return settings, nil
}

type YourPayloadStruct struct {
	// Define your payload structure here
}

func NewAPIClient(settings *Settings) *APIClient {
	authString := fmt.Sprintf("%s:%s", settings.Username, settings.Password)
	authHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(authString))

	client := resty.New()
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		req.Header.Set(authenticationHeader, authHeaderValue)
		return nil
	})

	client.BaseURL = settings.BaseURL

	if settings.ProxyUrl != "" {
		client.SetProxy(settings.ProxyUrl)
	}

	return &APIClient{client: client}
}

func (apiClient *APIClient) Get(endpoint string) (*resty.Response, error) {
	log.Println("GET: ", apiClient.client.BaseURL+endpoint)
	return apiClient.client.R().Get(apiClient.client.BaseURL + endpoint)
}

func (apiClient *APIClient) Post(endpoint string, payload interface{}) (*resty.Response, error) {
	log.Println("POST: ", apiClient.client.BaseURL+endpoint)
	return apiClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(apiClient.client.BaseURL + endpoint)
}

func (apiClient *APIClient) Put(endpoint string, payload interface{}) (*resty.Response, error) {
	log.Println("PUT: ", apiClient.client.BaseURL+endpoint)
	return apiClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Put(apiClient.client.BaseURL + endpoint)
}
