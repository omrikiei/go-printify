package go_printify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	contentType = "application/json;charset=utf-8"
	baseURL     = "api.printify.com"
	scheme      = "https"
)

type ApiRequest interface {
	GetMethod() string
	GetPath() string
	GetBody() interface{}
	GetResponseStruct() *interface{}
}

type Client struct {
	BaseURL    *url.URL
	ApiVersion string
	UserAgent  string
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: &url.URL{
			Scheme: scheme,
			Host:   baseURL,
		},
		UserAgent:  "go-printify v1.0",
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
	}
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: fmt.Sprintf("%s/%s", c.ApiVersion, path)}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode >= 400 {
		return resp, errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
