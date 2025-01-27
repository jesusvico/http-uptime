package endpoint

import (
	"fmt"
	"net/http"
	"net/url"
)

type Endpoint struct {
	Name   string
	Url    url.URL
	Method string
}

var validSchemes = map[string]bool{
	"http":  true,
	"https": true,
}

var validMethods = map[string]bool{
	http.MethodGet:    true,
	http.MethodPost:   true,
	http.MethodDelete: true,
	http.MethodPut:    true,
}

func New(name, urlStr, method string) (*Endpoint, error) {
	var e Endpoint

	e.Name = name

	// Validate the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL %s for endpoint %s: %v", urlStr, name, err)
	}
	if !validSchemes[parsedURL.Scheme] {
		return nil, fmt.Errorf("Invalid scheme in URL %s for endpoint %s", urlStr, name)
	}
	e.Url = *parsedURL

	// Validate the method
	if method == "" {
		method = http.MethodGet // Default to GET
	} else if !validMethods[method] {
		return nil, fmt.Errorf("Invalid method %s for endpoint %s", method, name)
	}
	e.Method = method

	return &e, nil
}

func (e *Endpoint) Request() (*http.Response, error) {
	req, err := http.NewRequest(e.Method, e.Url.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
