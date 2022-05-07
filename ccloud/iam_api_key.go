package ccloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/electric-saw/ccloud-client-go/ccloud/common"
)

type CreateAPIKeyReq struct {
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Owner       struct {
		ID          string `json:"id"`
		Environment string `json:"environment"`
	} `json:"owner"`
	Resource struct {
		ID          string `json:"id"`
		Environment string `json:"environment"`
	} `json:"resource"`
}

type APIKey struct {
	common.BaseModel `json:",inline"`
	Spec             APIKeySpec `json:"spec"`
}

type APIKeySpec struct {
	Secret      string `json:"secret"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Owner       struct {
		ID           string `json:"id"`
		Environment  string `json:"environment"`
		Related      string `json:"related"`
		ResourceName string `json:"resource_name"`
		APIVersion   string `json:"api_version"`
		Kind         string `json:"kind"`
	} `json:"owner"`
	Resource struct {
		ID           string `json:"id"`
		Environment  string `json:"environment"`
		Related      string `json:"related"`
		ResourceName string `json:"resource_name"`
		APIVersion   string `json:"api_version"`
		Kind         string `json:"kind"`
	} `json:"resource"`
}

func (c *ConfluentClient) CreateAPIKey(create *CreateAPIKeyReq) (*APIKey, error) {
	urlPath := "/iam/v2/api-keys"
	req, err := c.doRequest(urlPath, http.MethodPost, specWrap{create}, nil)
	if err != nil {
		return nil, err
	}

	if http.StatusAccepted != req.StatusCode {
		body, _ := ioutil.ReadAll(req.Body)
		log.Print(string(body))
		return nil, fmt.Errorf("failed to create api key: %s", req.Status)
	}
	var apiKey APIKey
	err = json.NewDecoder(req.Body).Decode(&apiKey)
	if err != nil {
		return nil, err
	}

	return &apiKey, nil
}
