package base

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type AIService struct {
	apiKey   string
	Endpoint string
	Model    string
	client   *resty.Client
}

func NewAIService(apiKey, endpoint, model string) *AIService {
	client := resty.New().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey)

	return &AIService{
		apiKey:   apiKey,
		Endpoint: endpoint,
		client:   client,
		Model:    model,
	}
}

func (s *AIService) Request(ctx context.Context, request AIRequest) (*AIResponse, error) {
	fullRequest := FullAIRequest{
		Model:     s.Model,
		AIRequest: request,
	}

	var response AIResponse

	responseRaw, err := s.client.R().
		SetContext(ctx).
		SetBody(fullRequest).
		SetDoNotParseResponse(true).
		Post(s.Endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to make ai request: %w", err)
	}
	defer responseRaw.RawBody().Close()

	decoder := json.NewDecoder(responseRaw.RawBody())
	if err := decoder.Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode ai response: %w", err)
	}

	return &response, nil
}
