package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/aejoy/vkgo/update"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetAlbums(properties ...any) ([]responses.Album, error) {
	query, reply := make(url.Values), responses.AlbumsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:  schema.NewType(schema.ParameterNames{"owner_id"}),
		schema.String:   schema.NewType(schema.ParameterNames{"action"}),
		schema.Duration: schema.NewType(schema.ParameterNames{"for"}),
		schema.Struct:   nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Albums, api.Call("photos.getAlbums", query.Encode(), &reply)
}

func (api *API) GetPhotos(properties ...any) ([]update.Photo, error) {
	query, reply := make(url.Values), responses.AlbumPhotosReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"owner_id", "album_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"action"}),
		schema.Boolean: schema.NewType(schema.ParameterNames{"rev"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Photos, api.Call("photos.get", query.Encode(), &reply)
}
