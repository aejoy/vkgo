package longpoll

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aejoy/vkgo/consts"
)

func Start(properties ...any) error {
	lp, err := New(properties...)
	if err != nil {
		return err
	}

	return lp.Start()
}

func (lp *LongPoll) Start() error {
	if lp.Bot == nil {
		return ErrSessionClosed
	}

	if lp.Session.Server == "" {
		return ErrSessionClosed
	}

	switch lp.BotType {
	case consts.CommunityBotType:
		lp.RunGroupLongPoll()
	case consts.UserBotType:
		lp.RunUserLongPoll()
	default:
		return nil
	}

	return nil
}

func (lp *LongPoll) getByteUpdates(server string) ([]byte, error) {
	link := fmt.Sprintf("%s?act=a_check&key=%s&ts=%+v&wait=%d&mode=%d&version=%d", server, lp.Session.Key, defineSessionTS(lp.Session.TS), lp.Timeout, lp.Mode, lp.Session.Version)

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)
	if err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll NewRequest error is %s", err.Error()))
		}

		return nil, err
	}

	response, err := lp.Bot.HTTPClient.Do(request)
	if err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll Do error is %s", err.Error()))
		}

		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		if err := response.Body.Close(); err != nil {
			if lp.Bot.Logger != nil {
				lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll BodyClose error is %s", err.Error()))
			}
		}

		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll ReadAll error is %s", err.Error()))
		}

		return nil, err
	}

	return body, nil
}

func defineSessionTS(sessionTS any) string {
	switch offset := sessionTS.(type) {
	case uint32, uint64:
		return fmt.Sprintf("%d", offset)
	case float32, float64:
		return fmt.Sprintf("%f", offset)
	default:
		return fmt.Sprint(offset)
	}
}
