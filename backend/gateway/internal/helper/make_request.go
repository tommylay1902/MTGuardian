package helper

import (
	"net/http"
	"strings"

	"github.com/tommylay1902/gateway/internal/constant"
)

func MakeRequest(method string, url string, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := http.Client{
		Timeout: constant.TIMEOUT,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// Handle error
		return nil, err
	}

	return resp, nil
}
