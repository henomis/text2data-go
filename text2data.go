package text2datago

import (
	"fmt"
	"time"

	"github.com/henomis/text2data-go/internal/pkg/httpclient"
	"github.com/henomis/text2data-go/pkg/request"
	"github.com/henomis/text2data-go/pkg/response"
)

const (
	Text2DataAPIEndpointV3 = "http://api.text2data.com/v3"
	analyzePath            = "/Analyze"
	categorizePath         = "/Categorize"
	extractPath            = "/Extract"
)

type Text2DataClient struct {
	privateKey string
	secret     string
	httpClient *httpclient.HttpClient
}

func New(endpoint, privateKey, secret string, timeout time.Duration) *Text2DataClient {
	return &Text2DataClient{
		privateKey: privateKey,
		secret:     secret,
		httpClient: httpclient.New(endpoint, timeout),
	}
}

func (m *Text2DataClient) Analyze(analyzeRequest *request.Request) (*response.Response, error) {

	analyzeRequest.PrivateKey = m.privateKey
	analyzeRequest.Secret = m.secret

	body, err := m.httpClient.Request(analyzePath, analyzeRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	analyzeResponse := &response.Response{}
	err = analyzeResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return analyzeResponse, nil
}

func (m *Text2DataClient) Categorize(categorizeRequest *request.Request) (*response.Response, error) {

	categorizeRequest.PrivateKey = m.privateKey
	categorizeRequest.Secret = m.secret

	body, err := m.httpClient.Request(categorizePath, categorizeRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	categorizeResponse := &response.Response{}
	err = categorizeResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return categorizeResponse, nil
}

func (m *Text2DataClient) Extract(extractRequest *request.Request) (*response.Response, error) {

	extractRequest.PrivateKey = m.privateKey
	extractRequest.Secret = m.secret

	body, err := m.httpClient.Request(extractPath, extractRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	extractResponse := &response.Response{}
	err = extractResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return extractResponse, nil
}
