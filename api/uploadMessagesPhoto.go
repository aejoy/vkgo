package api

import (
	"fmt"
	"net/url"

	"github.com/aejoy/vkgo/types"

	"github.com/aejoy/vkgo/responses"
)

func (api *API) GetUploadMessagesPhotoServer(chatID int) (server responses.UploadMessagesPhotoServer, err error) {
	reply := responses.UploadMessagesPhotoServerReply{}
	return reply.Response, api.Call("photos.getMessagesUploadServer", fmt.Sprintf("peer_id=%d", chatID), &reply)
}

func (api *API) SaveUploadMessagesPhoto(server int, photo, hash string) (photos []responses.SaveUploadMessagesPhoto, err error) {
	query, reply := url.Values{
		"server": {fmt.Sprint(server)},
		"photo":  {photo},
		"hash":   {hash},
	}, responses.SaveUploadMessagesPhotoReply{}

	return reply.Photos, api.Call("photos.saveMessagesPhoto", query.Encode(), &reply)
}

func (api *API) UploadMessagesPhoto(chatID int, fileName string, fileBytes []byte) (photo responses.SaveUploadMessagesPhoto, err error) {
	server, err := api.GetUploadMessagesPhotoServer(chatID)
	if err != nil {
		return photo, err
	}

	uploaded, err := api.UploadFiles(server.UploadURL, []types.UploadFile{
		{FieldName: "photo", FileName: fileName, Bytes: fileBytes},
	})
	if err != nil {
		return photo, err
	}

	photoFile := responses.UploadPhotoFile{}
	if err = uploaded.JSON(&photoFile); err != nil {
		return photo, err
	}

	photos, err := api.SaveUploadMessagesPhoto(photoFile.Server, photoFile.Photo, photoFile.Hash)
	if err != nil {
		return photo, err
	}

	if len(photos) > 0 {
		photo = photos[0]
	}

	return photo, err
}
