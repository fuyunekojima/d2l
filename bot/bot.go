package bot

import (
	"discord2line/line"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

var lineBot *line.Client

type DiscordBot struct {
	Session *discordgo.Session
}

func NewBot(token string) (*DiscordBot, error) {
	bot := DiscordBot{}
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	bot.Session = s
	bot.Session.AddHandler(onMessage)

	lineBot, err = line.NewLineClient(os.Getenv("LINE_SECRET"), os.Getenv("LINE_TOKEN"))
	if err != nil {
		return nil, err
	}
	return &bot, nil
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Message.Mentions) == 0 || len(m.Message.Mentions) >= 2 {
		return
	}
	botId := os.Getenv("BOT_ID")
	if m.Message.Mentions[0].ID != botId {
		return
	}
	recMes := regexp.MustCompile("<.*?>").ReplaceAllString(m.Message.Content, "")
	mes := fmt.Sprintf("User: %s\nMessage: \n%s", m.Message.Author.Username, recMes)
	err := lineBot.SendMessage(mes)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "メッセージ送信に失敗しました", m.MessageReference)
		return
	}
}

func (d *DiscordBot) Start() error {
	err := d.Session.Open()
	if err != nil {
		return err
	}
	defer d.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
	return nil
}

func (d *DiscordBot) Close() {
	d.Session.Close()
}
