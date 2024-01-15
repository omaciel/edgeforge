package clients

import (
	"encoding/base64"
	"fmt"
	"sync"

	"github.com/omaciel/edgeforge/config"
	log "github.com/sirupsen/logrus"

	"github.com/go-resty/resty/v2"
)

var (
	Client *APIClient
	lock   = &sync.Mutex{}
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

func Init() {
	Client = newAPIClient()
}

func Get() *APIClient {
	if Client == nil {
		lock.Lock()
		defer lock.Unlock()
		if Client == nil {
			Client = newAPIClient()
		}
	}
	return Client
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

func newAPIClient() *APIClient {
	cfg := config.Get()
	authString := fmt.Sprintf("%s:%s", cfg.Username, cfg.Password)
	authHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(authString))

	client := resty.New()
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		req.Header.Set(authenticationHeader, authHeaderValue)
		return nil
	})

	client.BaseURL = cfg.BaseURL

	if cfg.ProxyUrl != "" {
		client.SetProxy(cfg.ProxyUrl)
	}

	return &APIClient{client: client}
}

func (apiClient *APIClient) Get(endpoint string) (*resty.Response, error) {
	log.Debug("GET: ", apiClient.client.BaseURL+endpoint)
	return apiClient.client.R().Get(apiClient.client.BaseURL + endpoint)
}

func (apiClient *APIClient) Post(endpoint string, payload interface{}) (*resty.Response, error) {
	log.Debug("POST: ", apiClient.client.BaseURL+endpoint)
	return apiClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(apiClient.client.BaseURL + endpoint)
}

func (apiClient *APIClient) Put(endpoint string, payload interface{}) (*resty.Response, error) {
	log.Debug("PUT: ", apiClient.client.BaseURL+endpoint)
	return apiClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Put(apiClient.client.BaseURL + endpoint)
}
