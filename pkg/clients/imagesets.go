package clients

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetImageSetsList returns a list of image sets
func (apiClient *APIClient) GetImageSetsList() (*resty.Response, error) {
	endpoint := "/image-sets/view"
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageSetsImages(imagesetId int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/image-sets/view/%v/versions", imagesetId)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageSetsView(id int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/image-sets/view/%v/", id)
	return apiClient.Get(endpoint)
}

func (apiClient *APIClient) GetImageSetsImageView(imagesetId, imageVersion int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/image-sets/view/%v/versions/%v", imagesetId, imageVersion)
	return apiClient.Get(endpoint)
}
