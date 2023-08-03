package yandex300

import "net/http"

const (
	yandex300ApiUrl = "https://300.ya.ru/api/sharing-url"
)

type ClientConfig struct {
	authToken string

	BaseUrl    string
	HTTPClient *http.Client
}

func GetConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken:  authToken,
		BaseUrl:    yandex300ApiUrl,
		HTTPClient: &http.Client{},
	}
}
