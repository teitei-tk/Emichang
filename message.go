package emichang

import (
	"github.com/teitei-tk/emichang/line"
	"github.com/teitei-tk/emichang/util"

	"github.com/line/line-bot-sdk-go/linebot"
)

func HandleMessage(client *line.Client, event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		handleTextMessage(client, message, event)
	}
	return
}

func handleTextMessage(client *line.Client, message *linebot.TextMessage, event *linebot.Event) {
	replySource := util.GetReplySource(event)
	if message.Text == "Emichang ping" {
		textMsg := linebot.NewTextMessage("Emichangです！")
		client.Bot.PushMessage(replySource, textMsg)
	}
}
