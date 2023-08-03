package summarizer

import (
	"context"
	"sync"

	"github.com/Alexandr-Penkin/yandex300-summarizer/internal/yandex300"
)

type Summarizer interface {
	GetSummary(cnt context.Context, url string) (string, error)
}

type YandexSummarizer struct {
	client *yandex300.Client
	mu     sync.Mutex
}

func New(oAuthToken string) *YandexSummarizer {
	s := &YandexSummarizer{
		client: yandex300.NewClient(oAuthToken),
	}

	return s
}

func (s *YandexSummarizer) GetSummary(ctx context.Context, url string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	request, err := s.client.NewRequest(ctx, url)
	if err != nil {
		return "", err
	}

	sharingUrl, err := s.client.SendRequest(request)
	if err != nil {
		return "", err
	}

	return sharingUrl, nil
}
