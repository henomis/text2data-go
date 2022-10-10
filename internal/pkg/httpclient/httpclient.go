package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	httpClient *http.Client
	baseURL    string
}

type RequestData interface {
	ToJSON() ([]byte, error)
}

func New(baseURL string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

func (h *HttpClient) Request(path string, requestData RequestData) (io.ReadCloser, error) {

	jsonData, err := requestData.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to convert request data to json: %w", err)
	}

	request, err := http.NewRequest("POST", h.baseURL+path, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", response.StatusCode)
	}

	return response.Body, nil
}
