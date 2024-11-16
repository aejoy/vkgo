package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) AddLike(properties ...any) (likes int, reactions responses.Reactions, err error) {
	query, reply := make(url.Values), responses.LikeActionReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"owner_id", "item_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type", "access_key"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Likes, reply.Response.Reactions, api.Call("likes.add", query.Encode(), &reply)
}

func (api *API) DeleteLike(properties ...any) (likes int, reactions responses.Reactions, err error) {
	query, reply := make(url.Values), responses.LikeActionReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"owner_id", "item_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type", "access_key"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Likes, reply.Response.Reactions, api.Call("likes.delete", query.Encode(), &reply)
}

func (api *API) GetLikes(properties ...any) (likes []responses.User, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.LikesReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"owner_id", "item_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type", "filter"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Users, api.Call("likes.getList", query.Encode(), &reply)
}

func (api *API) IsLiked(properties ...any) (liked bool, reactionID int, copied bool, err error) {
	query, reply := make(url.Values), responses.IsLikedReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"owner_id", "item_id", "user_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Liked == 1, reply.Response.ReactionID, reply.Response.Copied == 1, api.Call("likes.isLiked", query.Encode(), &reply)
}
