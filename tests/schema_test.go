package tests_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/botscommunity/botsgo/pkg/schema"
)

func TestSchema(t *testing.T) {
	t.Run("typeDefs", func(t *testing.T) {
		query := make(url.Values)

		schema.NewSchema(schema.TypeDefs{
			schema.Integer:      schema.NewType(schema.ParameterNames{"chat_ids", "chat_message_id", "offset", "count"}),
			schema.String:       schema.NewType(schema.ParameterNames{"text"}),
			schema.ArrayString:  schema.NewType(schema.ParameterNames{"attachments"}),
			schema.ArrayInteger: schema.NewType(schema.ParameterNames{"chat_ids"}),
			schema.Struct:       nil,
		}).ConvertToQuery(query, []int{1, 2, 3}, 1, struct {
			Text string `json:"text"`
		}{Text: "meow :3"})

		assert.Equal(t, true, query.Encode() == "chat_ids=1%2C2%2C3&chat_message_id=1&text=meow+%3A3", "Schema is broken")
	})
}
