package clients

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	models "github.com/omaciel/edgeforge/pkg/models/images"
)

const (
	baseURLEnv           = "API_BASEURL"
	usernameEnv          = "API_USERNAME"
	passwordEnv          = "API_PASSWORD"
	proxyHTTP            = "API_HTTP_PROXY"
	authenticationHeader = "Authorization"
)

var (
	ErrNoUsernameProvided = errors.New("no username provided")
	ErrNoPasswordProvided = errors.New("no password provided")
	ErrNoBaseUrlProvided  = errors.New("no baseURL provided")
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
	return apiClient.client.R().Get(apiClient.client.BaseURL + endpoint)
}

func (apiClient *APIClient) Post(endpoint string, payload interface{}) (*resty.Response, error) {
	return apiClient.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(apiClient.client.HostURL + endpoint)
}

func (apiClient *APIClient) CreateImage(image *models.Image) (*resty.Response, error) {
	return apiClient.Post("/images", image)
}

func (apiClient *APIClient) GetImageSetView(id int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/image-sets/view/%v/", id)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageSetImageView(imagesetId, imageVersion int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/image-sets/view/%v/versions/%v", imagesetId, imageVersion)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageDetails(imageVersion int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/images/%v/details", imageVersion)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageRepo(imageVersion int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/images/%v/repo", imageVersion)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageStatus(imageVersion int) (*resty.Response, error) {
	return apiClient.Get(fmt.Sprintf("/images/%v/status", imageVersion))
}

func (apiClient *APIClient) UpdateImage(imageVersion int, image *models.Image) (*resty.Response, error) {
	return apiClient.Post(fmt.Sprintf("/images/%v/update", imageVersion), image)
}
