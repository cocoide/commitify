package gateway

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type HttpClient struct {
	client   *http.Client
	endpoint string
	headers  map[string]string
	params   map[string]interface{}
	body     io.Reader
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client:  &http.Client{},
		headers: make(map[string]string),
		params:  make(map[string]interface{}),
	}
}

func (h *HttpClient) WithBaseURL(baseURL string) *HttpClient {
	h.endpoint = baseURL
	return h
}

func (h *HttpClient) WithHeader(key, value string) *HttpClient {
	h.Headers[key] = value
	return h
}

func (h *HttpClient) WithBearerToken(token string) *HttpClient {
	h.headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return h
}

func (h *HttpClient) WithPath(path string) *HttpClient {
	h.endpoint = h.endpoint + "/" + path
	return h
}

func (h *HttpClient) WithParam(key string, value interface{}) *HttpClient {
	h.params[key] = value
	return h
}

type HttpMethod int

const (
	GET HttpMethod = iota + 1
	POST
	DELTE
	PUT
)

func (h *HttpClient) WithBody(values []byte) *HttpClient {
	h.Body = bytes.NewReader(values)
	return h
}

func (h *HttpClient) Execute(method HttpMethod) ([]byte, error) {
	var methodName string
	switch method {
	case GET:
		methodName = "GET"
	case POST:
		methodName = "POST"
	case DELTE:
		methodName = "DELETE"
	case PUT:
		methodName = "PUT"
	}
	client := h.client

	req, err := http.NewRequest(methodName, h.endpoint, h.body)
	if err != nil {
		return nil, err
	}

	for k, v := range h.headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range h.params {
		switch v := value.(type) {
		case string:
			query.Add(key, v)
		case int:
			query.Add(key, strconv.Itoa(v))
		case bool:
			query.Add(key, strconv.FormatBool(v))
		default:
			return nil, fmt.Errorf("Failed to parse param value: %v", value)
		}
	}
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
