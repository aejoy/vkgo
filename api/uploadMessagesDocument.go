package api

import (
	"net/url"

	"github.com/aejoy/vkgo/types"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetUploadMessagesDocumentServer(properties ...any) (server responses.UploadMessagesDocumentServer, err error) {
	query, reply := make(url.Values), responses.UploadMessagesDocumentServerReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("docs.getMessagesUploadServer", query.Encode(), &reply)
}

func (api *API) SaveUploadMessagesDocument(file string) (typ string, document responses.SaveUploadMessageDocument, err error) {
	query, reply := url.Values{
		"file": {file},
	}, responses.SaveUploadMessagesDocumentReply{}

	return reply.Response.Type, reply.Response.Document, api.Call("docs.save", query.Encode(), &reply)
}

func (api *API) UploadMessagesDocument(chatID int, file types.UploadFile) (typ string, document responses.SaveUploadMessageDocument, err error) {
	server, err := api.GetUploadMessagesDocumentServer(chatID)
	if err != nil {
		return "", document, err
	}

	uploaded, err := api.UploadFiles(server.UploadURL, []types.UploadFile{file})
	if err != nil {
		return "", document, err
	}

	documentFile := responses.UploadDocumentFile{}
	if err = uploaded.JSON(&documentFile); err != nil {
		return "", document, err
	}

	return api.SaveUploadMessagesDocument(documentFile.File)
}
