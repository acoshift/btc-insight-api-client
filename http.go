package btcinsightapiclient

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// Errors
var (
	ErrAPIError = errors.New("api error")
)

func invokeHTTP(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrAPIError
	}
	return buf.Bytes(), nil
}
