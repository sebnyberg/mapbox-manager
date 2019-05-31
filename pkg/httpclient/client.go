package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request http.Request

type APIResponse struct {
	Payload    []byte
	StatusCode int
}

// RequestTransform performs some operation on the http request
// Transforms can be used to e.g. strip the request of a single header or query param
// allowing for easy testing of responses due to missing params
type RequestTransform func(req *http.Request) *http.Request

type Client struct {
	client *http.Client

	// Transforms are run after method execution
	transforms  []RequestTransform
	headers     map[string]string
	queryParams map[string]interface{}
	baseURL     string
}

type Config struct {
	DefaultQueryParams map[string]interface{}
	DefaultHeaders     map[string]string
	BaseURL            string
}

func NewClient() Client {
	return Client{
		client:      &http.Client{},
		headers:     make(map[string]string),
		queryParams: make(map[string]interface{}),
		baseURL:     "",
	}
}

// Returns a copy of the client with headers, params and baseURL set by the config
func (c Client) WithConfig(conf *Config) Client {
	if conf.DefaultHeaders != nil {
		c.headers = conf.DefaultHeaders
	}
	if conf.DefaultQueryParams != nil {
		c.queryParams = conf.DefaultQueryParams
	}
	c.baseURL = conf.BaseURL

	return c
}

// Add a header to the list of default headers
func (c Client) WithHeader(key, value string) Client {
	c.headers[key] = value

	return c
}

// Add headers to the list of default headers
func (c Client) WithHeaders(headers map[string]string) Client {
	for k, v := range headers {
		c.headers[k] = v
	}

	return c
}

// Ensure that a header is removed before requests are performed
func (c Client) EnsureWithoutHeader(headerName string) Client {
	c.transforms = append(c.transforms, func(r *http.Request) *http.Request {
		r.Header.Del(headerName)

		return r
	})

	return c
}

// Ensure that a query parameter is removed before requests are performed
func (c Client) EnsureWithoutQueryParam(paramName string) Client {
	c.transforms = append(c.transforms, func(r *http.Request) *http.Request {
		queryParams := r.URL.Query()

		queryParams.Del(paramName)

		r.URL.RawQuery = queryParams.Encode()

		return r
	})

	return c
}

// Add query parameters to default query parameters
func (c Client) WithQueryParams(newParams map[string]interface{}) Client {
	for k, v := range newParams {
		c.queryParams[k] = v.(string)
	}

	return c
}

// Performs a request
func (c *Client) performRequest(
	r *http.Request,
	headers map[string]string,
	params map[string]interface{}) (*APIResponse, error) {

	var req *Request = (*Request)(r)

	// Add headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Add params
	queryParams := r.URL.Query()
	for k, v := range c.queryParams {
		queryParams.Add(k, fmt.Sprintf("%v", v))
	}
	for k, v := range params {
		queryParams.Add(k, fmt.Sprintf("%v", v))
	}
	r.URL.RawQuery = queryParams.Encode()

	// Run any transforms which are pending
	for _, transform := range c.transforms {
		r = transform(r)
	}

	// Perform request
	resp, err := c.client.Do(r)
	if err != nil {
		log.Fatalln(err)
	}

	// Load response
	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return &APIResponse{
		respBytes,
		resp.StatusCode,
	}, nil
}

// Attempt at copying the behaviour of axios.get (JS) / resquests.get (Python)
func (c *Client) Get(url string, headers map[string]string, params map[string]interface{}) (*APIResponse, error) {
	finalURL := c.baseURL + url
	r, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}

	return c.performRequest(r, headers, params)
}

func (c *Client) Delete(url string, headers map[string]string, params map[string]interface{}) (*APIResponse, error) {
	finalURL := c.baseURL + url

	r, err := http.NewRequest("DELETE", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}

	return c.performRequest(r, headers, params)
}

func (c *Client) Post(url string, headers map[string]string, payload map[string]interface{}) (*APIResponse, error) {
	finalURL := c.baseURL + url

	if payload == nil {
		return nil, errors.New("Cannot POST without a payload")
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse request payload, error: %v", err)
	}

	payloadBuffer := bytes.NewBuffer(payloadBytes)

	r, err := http.NewRequest("POST", finalURL, payloadBuffer)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}

	return c.performRequest(r, headers, nil)
}
