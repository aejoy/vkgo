package responses

type UploadPhotoFile struct {
	Server     int    `json:"server,omitempty"`
	PhotosList string `json:"photos_list,omitempty"`
	Photo      string `json:"photo,omitempty"`
	AlbumID    int    `json:"aid,omitempty"`
	Hash       string `json:"hash,omitempty"`
}

type UploadDocumentFile struct {
	File string `json:"file,omitempty"`
}
