package src

import (
	"context"
	"net/http"
	"time"
)

const GraphBaseURL string = "https://graph.microsoft.com/v1.0/"

func NewGraphClient(apiKey string) *Client {

	return &Client{
		BaseURL: GraphBaseURL,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c Client) Request(ctx context.Context, method string, request string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(
		method,
		c.BaseURL+request,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}
