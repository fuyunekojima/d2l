package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Client struct {
	Bot *linebot.Client
}

func NewLineClient(secret, token string) (*Client, error) {
	bot, err := linebot.New(secret, token)
	if err != nil {
		return nil, err
	}
	c := Client{
		Bot: bot,
	}
	return &c, nil
}

func (c *Client) SendMessage(message string) error {
	var messages []linebot.SendingMessage = make([]linebot.SendingMessage, 0)
	messages = append(messages, linebot.NewTextMessage(message))
	_, err := c.Bot.BroadcastMessage(messages...).Do()
	if err != nil {
		return err
	}
	return nil
}
