package api

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/aejoy/vkgo/responses"
	"github.com/aejoy/vkgo/types"
)

func (api *API) UploadFiles(serverURL string, files []types.UploadFile) (uploadedFile responses.UploadedFiles, err error) {
	var (
		body  = new(bytes.Buffer)
		write = multipart.NewWriter(body)
	)

	for _, file := range files {
		part, err := write.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, bytes.NewReader(file.Bytes))
		if err != nil {
			return nil, err
		}
	}

	contentType := write.FormDataContentType()
	_ = write.Close()

	request, err := http.NewRequestWithContext(context.Background(), "POST", serverURL, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", contentType)

	res, err := api.Client.HTTPClient.Do(request)
	if err != nil {
		err := res.Body.Close()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
