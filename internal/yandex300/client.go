package yandex300

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	config ClientConfig
}

func NewClient(authTocken string) *Client {
	config := GetConfig(authTocken)

	return &Client{
		config: config,
	}
}

type Request struct {
	ArticleUrl string `json:"article_url"`
}

type Response struct {
	Status     string `json:"status"`
	Message    string `json:"message,omitempty"`
	SharingUrl string `json:"sharing_url,omitempty"`
}

func (c *Client) NewRequest(ctx context.Context, url string) (*http.Request, error) {
	requestBody := Request{
		ArticleUrl: url,
	}

	marshaled, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.config.BaseUrl, bytes.NewReader(marshaled))
	if err != nil {
		return nil, err
	}

	var oAuth = "OAuth " + c.config.authToken

	req.Header.Set("Authorization", oAuth)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) SendRequest(req *http.Request) (string, error) {
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("Bad response code: %d" + resp.Status)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Status != "success" {
		return "", errors.New(response.Message)
	}

	return response.SharingUrl, err
}
