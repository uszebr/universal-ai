package base

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uszebr/universal-ai/util"
)

func TestAIService_Request(t *testing.T) {
	// Mock OpenAI API response
	mockResponse := `{
		"id": "chatcmpl-123",
		"object": "chat.completion",
		"created": 1700000000,
		"choices": [{
			"index": 0,
			"message": {
				"role": "assistant",
				"content": "Hello, how can I assist you?"
			},
			"finish_reason": "stop"
		}]
	}`

	// Create a test server to mock OpenAI API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Contains(t, r.Header.Get("Authorization"), "Bearer ")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(mockResponse))
		assert.NoError(t, err)
	}))
	defer server.Close()

	service := NewAIService("test-api-key", server.URL, "test-model")
	req := AIRequest{
		Messages: []Message{
			{Role: UserRole, Content: util.StrPtr("test message")},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := service.Request(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.Choices, "Choices should not be empty")
	assert.Equal(t, "chatcmpl-123", resp.ID)
	assert.Equal(t, "chat.completion", resp.Object)
	assert.Len(t, resp.Choices, 1)
	assert.Equal(t, AssistantRole, resp.Choices[0].Message.Role)
	assert.Equal(t, "Hello, how can I assist you?", *resp.Choices[0].Message.Content)
	assert.Equal(t, "stop", resp.Choices[0].FinishReason)
}

func TestAIService_Request_WithTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate delay
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := NewAIService("test-api-key", server.URL, "test-model")

	req := AIRequest{
		Messages: []Message{
			{Role: UserRole, Content: util.StrPtr("test message")},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	resp, err := service.Request(ctx, req)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestAIService_Request_WithCancelledContext(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate delay
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := NewAIService("test-api-key", server.URL, "test-model")

	req := AIRequest{
		Messages: []Message{
			{Role: UserRole, Content: util.StrPtr("test message")},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Immediately cancel the context

	// Make the request with the cancelled context
	resp, err := service.Request(ctx, req)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "context canceled")
}
