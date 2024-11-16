<img src="preview.png" alt="Bots Community">

<p align="center">Flexible and high-performance VK API module Go</p>

<div align="center">
	<a href="https://pkg.go.dev/github.com/aejoy/vkgo">
		<img src="https://img.shields.io/static/v1?label=go&message=reference&color=00add8&logo=go" />
	</a>
	<a href="http://www.opensource.org/licenses/MIT">
		<img src="https://img.shields.io/badge/license-MIT-fuchsia.svg" />
	</a>
    <a href="https://goreportcard.com/report/github.com/aejoy/vkgo">
		<img src="https://goreportcard.com/badge/github.com/aejoy/vkgo" />
	</a>
</div>

<h2 align="center">Instalation</h2>

```bash
go get github.com/aejoy/vkgo
```

<h2 align="center">Getting started</h2>

Examples of working bots can be seen in the catalog [/samples](./samples)

A simple example of a LongPoll API Bot:

```go
package main

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/scene"
	"github.com/aejoy/vkgo/update"
	"log"
	"os"
)

func main() {
	bot, err := api.NewBot(os.Getenv("TOKEN"))
	if err != nil { /* ... */
	}

	messages := scene.OnMessage(func(bot *api.API, message update.Message) {
		bot.SendMessage(message.ChatID, message.Text)
	})

	log.Fatalln(longpoll.Start(bot, messages))
}

```

<h2 align="center">Help in solving problems</h2>

Don't know how to solve your problem? Ask the programmers from [our community](./CONTACTS.md). There is a chance that they have already dealt with this problem and are ready to help you