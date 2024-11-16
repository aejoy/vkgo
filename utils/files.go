package utils

import (
	"context"
	"io"
	"net/http"
)

func GetFileFromURL(url string) (data []byte, err error) {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data = body

	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return
}
