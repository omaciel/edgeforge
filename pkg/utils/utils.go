package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func readSettingsFromFile(filename string) (*struct{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var respObj struct{}
	if err := yaml.Unmarshal(data, &respObj); err != nil {
		return nil, err
	}

	return &respObj, nil
}
