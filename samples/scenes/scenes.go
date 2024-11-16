package scenes

import (
	"fmt"
	"time"

	"github.com/aejoy/vkgo/utils"

	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/types"
	"github.com/aejoy/vkgo/update"
)

func NewMessageScene(bot *api.API, message update.Message) {
	switch message.Text {
	case ":пинг":
		text, startTime := "✦ Понг!", time.Now()

		sent, err := bot.SendMessage(message.ChatID, text)
		if err != nil {
			panic(err)
		}

		if _, err := bot.EditMessage(sent.ChatID, sent.ChatMessageID,
			fmt.Sprintf("%s (%v)",
				text,
				time.Since(startTime),
			),
		); err != nil {
			panic(err)
		}
	case ":upload":
		fileByte, err := utils.GetFileFromURL("https://4kwallpapers.com/images/walls/thumbs_3t/18949.jpg")
		if err != nil {
			panic(err)
		}

		photo, err := bot.UploadMessagesPhoto(message.ChatID, "wallpaper-4k.png", fileByte)
		if err != nil {
			panic(err)
		}

		if _, err := bot.SendMessage(types.SendMessage{
			ChatID:     message.ChatID,
			Attachment: fmt.Sprintf("photo%d_%d_%s", photo.OwnerID, photo.ID, photo.AccessKey),
		}); err != nil {
			panic(err)
		}
	case ":album upload":
		fileByte, err := utils.GetFileFromURL("https://i.pinimg.com/736x/4d/1d/b0/4d1db0811c9c0f48cf7ed11c6ab00a7b.jpg")
		if err != nil {
			panic(err)
		}

		clubAlbumID, groupID := 301684048, 226420674

		photos, err := bot.UploadAlbumPhotos(clubAlbumID, []types.UploadFile{
			{FieldName: "file1", FileName: "meme.png", Bytes: fileByte},
		}, groupID)
		if err != nil {
			panic(err)
		}

		if _, err := bot.SendMessage(types.SendMessage{
			ChatID:     message.ChatID,
			Attachment: fmt.Sprintf("photo%d_%d", photos[0].OwnerID, photos[0].ID),
		}); err != nil {
			panic(err)
		}
	}
}
