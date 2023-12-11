package gateway

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type HttpClient struct {
	Client   *http.Client
	Endpoint string
	Headers  map[string]string
	Params   map[string]interface{}
	Body     io.Reader
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client:  &http.Client{},
		Headers: make(map[string]string),
		Params:  make(map[string]interface{}),
	}
}

func (h *HttpClient) WithBaseURL(baseURL string) *HttpClient {
	h.Endpoint = baseURL
	return h
}

func (h *HttpClient) WithHeader(key, value string) *HttpClient {
	h.Headers[key] = value
	return h
}

func (h *HttpClient) WithBearerToken(token string) *HttpClient {
	h.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return h
}

func (h *HttpClient) WithPath(path string) *HttpClient {
	h.Endpoint = h.Endpoint + "/" + path
	return h
}

func (h *HttpClient) WithParam(key string, value interface{}) *HttpClient {
	h.Params[key] = value
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
	client := h.Client

	req, err := http.NewRequest(methodName, h.Endpoint, h.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range h.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range h.Params {
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
