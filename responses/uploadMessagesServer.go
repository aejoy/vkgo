package responses

type UploadMessagesPhotoServerReply struct {
	ErrorInterface
	Response UploadMessagesPhotoServer `json:"response,omitempty"`
}

type UploadMessagesPhotoServer struct {
	AlbumID   int    `json:"album_id,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	UploadURL string `json:"upload_url,omitempty"`
}

type UploadMessagesDocumentServerReply struct {
	ErrorInterface
	Response UploadMessagesDocumentServer `json:"response,omitempty"`
}

type UploadMessagesDocumentServer struct {
	UploadURL string `json:"upload_url,omitempty"`
}
