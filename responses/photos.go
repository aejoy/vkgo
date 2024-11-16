package responses

import "github.com/aejoy/vkgo/update"

type AlbumPhotosReply struct {
	ErrorInterface
	Response AlbumPhotos `json:"response"`
}

type AlbumPhotos struct {
	Photos []update.Photo `json:"items"`
}
