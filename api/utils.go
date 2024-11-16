package api

import (
	"fmt"

	"github.com/aejoy/vkgo/responses"
)

func (api *API) ResolveDomain(domain string) (typ string, objectID int, err error) {
	reply := responses.ResolveDomainReply{}
	return reply.Response.Type, reply.Response.ObjectID, api.Call("utils.resolveScreenName", fmt.Sprintf("screen_name=%s", domain), &reply)
}
