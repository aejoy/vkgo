package tests_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aejoy/vkgo/api"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	t.Run("UserBot", func(t *testing.T) {
		fmt.Println("USER_TOKEN=", os.Getenv("USER_TOKEN"))

		_, err := api.New(os.Getenv("USER_TOKEN"))
		assert.Equal(t, true, err == nil, err)
	})

	t.Run("CommunityBot", func(t *testing.T) {
		fmt.Println("GROUP_TOKEN=", os.Getenv("GROUP_TOKEN"))

		_, err := api.New(os.Getenv("GROUP_TOKEN"))
		assert.Equal(t, true, err == nil, err)
	})
}
