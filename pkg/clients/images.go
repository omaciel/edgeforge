package clients

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/omaciel/edgeforge/pkg/types"
)

func (apiClient *APIClient) CreateImage(image *types.Image) (*resty.Response, error) {
	endpoint := "/images"
	return apiClient.Post(endpoint, image)
}

func (apiClient *APIClient) GetImageDetails(imageVersion int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/images/%v", imageVersion)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageRepo(imageVersion int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/images/%v/repo", imageVersion)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageStatus(imageVersion int) (*resty.Response, error) {
	return apiClient.Get(fmt.Sprintf("/images/%v/status", imageVersion))
}

func (apiClient *APIClient) UpdateImage(imageVersion int, image *types.Image) (*resty.Response, error) {
	return apiClient.Post(fmt.Sprintf("/images/%v/update", imageVersion), image)
}
