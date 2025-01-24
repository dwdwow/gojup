package gojup

import (
	"fmt"
	"io"
	"net/http"
)

func simpleGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("jupiter: GET %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("jupiter: GET %s: %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("jupiter: GET %s: %w", url, err)
	}
	return body, nil
}
