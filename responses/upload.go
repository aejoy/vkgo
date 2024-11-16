package responses

import "encoding/json"

type UploadedFiles []byte

func (uploadedFiles UploadedFiles) JSON(data any) error {
	return json.Unmarshal(uploadedFiles, data)
}
