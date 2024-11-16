package benchmarks_test

import (
	"os"
	"strconv"
	"testing"

	// "go.uber.org/zap".

	vksdkApi "github.com/SevereCloud/vksdk/api"
	vkgoApi "github.com/aejoy/vkgo/api"
)

func Benchmark_VKGO(b *testing.B) {
	communityID, err := strconv.Atoi(os.Getenv("COMMUNITY_ID"))
	if err != nil {
		b.Fatal(err)
	}

	bot, err := vkgoApi.NewCommunity(communityID, os.Getenv("TOKEN"))
	if err != nil {
		b.Fatal(err)
	}

	b.Run("getUser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := bot.GetUser(542439242); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_VKSDK(b *testing.B) {
	bot := vksdkApi.NewVK(os.Getenv("TOKEN"))

	b.Run("getUser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := bot.UsersGet(vksdkApi.Params{
				"user_ids": 542439242,
			}); err != nil {
				b.Fatal(err)
			}
		}
	})
}
