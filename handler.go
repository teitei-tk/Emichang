package emichang

import (
	"log"
	"net/http"

	"github.com/teitei-tk/emichang/line"

	"github.com/line/line-bot-sdk-go/linebot"
)

// line api callback handler func.
func HandleFunc(w http.ResponseWriter, req *http.Request) {
	client, err := line.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	events, err := client.Bot.ParseRequest(req)
	if err != nil {
		client.Logger.Fatal(err)
		invalidLineResponse(w, err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			HandleMessage(client, event)
		}
	}

	return
}

func invalidLineResponse(w http.ResponseWriter, err error) {
	if err == linebot.ErrInvalidSignature {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(500)
	}

	return
}
