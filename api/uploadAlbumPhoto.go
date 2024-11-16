package api

import (
	"net/url"

	"github.com/aejoy/vkgo/types"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetUploadAlbumPhotoServer(properties ...any) (server responses.UploadMessagesPhotoServer, err error) {
	query, reply := make(url.Values), responses.UploadMessagesPhotoServerReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"album_id", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("photos.getUploadServer", query.Encode(), &reply)
}

func (api *API) SaveUploadAlbumPhoto(properties ...any) (photos []responses.SaveUploadAlbumPhoto, err error) {
	query, reply := make(url.Values), responses.SaveUploadAlbumPhotoReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"server", "album_id", "group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"photos_list", "hash", "caption"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("photos.save", query.Encode(), &reply)
}

func (api *API) UploadAlbumPhotos(albumID int, files []types.UploadFile, properties ...any) (photo []responses.SaveUploadAlbumPhoto, err error) {
	args := []any{albumID}
	args = append(args, properties...)

	server, err := api.GetUploadAlbumPhotoServer(args...)
	if err != nil {
		return photo, err
	}

	uploaded, err := api.UploadFiles(server.UploadURL, files)
	if err != nil {
		return photo, err
	}

	photoFile := responses.UploadPhotoFile{}
	if err = uploaded.JSON(&photoFile); err != nil {
		return photo, err
	}

	saveArgs := []any{photoFile.Server}
	saveArgs = append(saveArgs, args...)
	saveArgs = append(saveArgs, photoFile.PhotosList, photoFile.Hash)

	photos, err := api.SaveUploadAlbumPhoto(saveArgs...)
	if err != nil {
		return photo, err
	}

	return photos, err
}
