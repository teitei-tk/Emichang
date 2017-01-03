package line

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Client is provide linebot api struct.
type Client struct {
	bot    *linebot.Client
	logger *log.Logger
}

// NewClient is provide line-bot api struct client.
func NewClient() (*Client, error) {
	channelSecret := os.Getenv("LINE_CHANNE_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")

	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	logger := log.New(ioutil.Discard, "", log.LstdFlags)

	return &Client{
		bot:    bot,
		logger: logger,
	}, nil
}
