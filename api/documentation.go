package api

import (
	"github.com/aejoy/vkgo/responses"
)

func (api *API) GetHealth() (statuses []responses.HealthStatus, err error) {
	reply := responses.HealthReply{}
	return statuses, api.Call("documentation.getPlatformHealthStatuses", "", &reply)
}
