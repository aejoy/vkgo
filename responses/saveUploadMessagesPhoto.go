package responses

import "github.com/aejoy/vkgo/update"

type SaveUploadMessagesPhotoReply struct {
	ErrorInterface
	Photos []SaveUploadMessagesPhoto `json:"response"`
}

type SaveUploadMessagesPhoto struct {
	Date      int    `json:"date,omitempty"`
	ID        int    `json:"id,omitempty"`
	AlbumID   int    `json:"album_id,omitempty"`
	OwnerID   int    `json:"owner_id,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
}

type SaveUploadAlbumPhotoReply struct {
	ErrorInterface
	Response []SaveUploadAlbumPhoto `json:"response"`
}

type SaveUploadAlbumPhoto struct {
	AlbumID int                `json:"album_id"`
	Date    int                `json:"date"`
	ID      int                `json:"id"`
	OwnerID int                `json:"owner_id"`
	Sizes   []update.PhotoSize `json:"sizes"`
	Text    string             `json:"text"`
	HasTags bool               `json:"has_tags"`
}
