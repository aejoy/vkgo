package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/aejoy/vkgo/consts"
	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo"
	"go.uber.org/zap"
)

type API struct {
	*botsgo.Client
	ID          int    `json:"id"`
	Token       string `json:"token"`
	Version     string `json:"version"`
	IsGroup     bool   `json:"is_group"`
	ContentType string
	Limit       int
	mutex       sync.Mutex
	time        time.Time
	rps         int
}

func NewCommunity(communityID int, token string, properties ...any) (*API, error) {
	return New(&API{
		ID:    communityID,
		Token: token,
		Limit: consts.CommunityAPILimit,
	}, properties)
}

func NewUser(userID int, token string, properties ...any) (*API, error) {
	return New(&API{
		ID:    userID,
		Token: token,
		Limit: consts.CommunityAPILimit,
	}, properties)
}

func New(properties ...any) (*API, error) {
	client, err := botsgo.NewClient("https://api.vk.com/method")
	if err != nil {
		return nil, err
	}

	client.Logger, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}

	bot := &API{
		Client:      client,
		Version:     consts.DefaultAPIVersion,
		ContentType: "application/msgpack",
		Limit:       consts.UserAPILimit,
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *API:
			botProperties(bot, property)
		case int:
			bot.ID = property
		case string:
			bot.Token = property
		case float32:
			bot.Version = fmt.Sprint(property)
		case *zap.Logger:
			bot.Client.Logger = property
		}
	}

	if bot.ID != 0 {
		return bot, nil
	}

	if group, err := bot.GetGroup(); err != nil {
		user, err := bot.GetUser()
		if err != nil {
			return nil, err
		}

		bot.ID = user.ID
	} else {
		bot.ID = group.ID
		bot.IsGroup = true
		bot.Limit = consts.CommunityAPILimit
	}

	bot.Client.Logger.Info(fmt.Sprintf("Connected to %d", bot.ID))

	return bot, nil
}

func botProperties(bot, property *API) {
	if property.ID != 0 {
		bot.ID = property.ID
	}

	if property.Token != "" {
		bot.Token = property.Token
	}

	if property.Version != "" {
		bot.Version = property.Version
	}

	if property.ContentType != "" {
		bot.ContentType = property.ContentType
	}

	if property.Limit != 0 {
		bot.Limit = property.Limit
	}

	if property.Client != nil {
		if property.Client.HTTPClient != nil {
			bot.Client.HTTPClient = property.Client.HTTPClient
		}

		if property.APIURL != "" {
			bot.Client.APIURL = property.APIURL
		}

		if property.Client.Logger != nil {
			bot.Client.Logger = property.Client.Logger
		}
	}
}

func (api *API) Call(method string, query string, response responses.Response) error {
	if api.Limit > 0 {
		api.mutex.Lock()

		sleepTime := time.Second - time.Since(api.time)
		if sleepTime < 0 {
			api.time = time.Now()
			api.rps = 0
		} else if api.rps == api.Limit {
			time.Sleep(sleepTime)
			api.time = time.Now()
			api.rps = 0
		}

		api.rps++
		api.mutex.Unlock()
	}

	// code from VK SDK, thanks <3

	if api.ContentType == "application/msgpack" {
		method += ".msgpack"
	}

	req, err := api.Client.NewRequest(context.Background())
	if err != nil {
		return err
	}

	req.Method(http.MethodGet)

	req.Path(fmt.Sprintf("/method/%s?v=%s&%s", method, api.Version, query))
	req.Response(&response)

	req.SetHeader("Authorization", "Bearer "+api.Token)

	res, err := req.Do()
	if err != nil {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return err
	}

	if err := response.GetError(); err != nil {
		return err
	}

	return nil
}
