package tgbot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

// CmdHandler func(Chat.ID, CommandArguments)
type CmdHandler func(string) string

type Commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]CmdHandler
}

func NewCommander() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Commander{
		bot:   bot,
		route: make(map[string]CmdHandler),
	}, nil
}

func (c *Commander) RegisterHandler(cmd string, f CmdHandler) {
	c.route[cmd] = f
}

func (c *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			if f, ok := c.route[cmd]; ok {
				msg.Text = f(update.Message.CommandArguments())
			} else {
				msg.Text = "Unknown command"
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf("you send <%v>", update.Message.Text)
		}
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
