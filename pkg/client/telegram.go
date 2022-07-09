package client

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramChatType int64

const (
	Channel TelegramChatType = 0
	Group                    = 1
)

var TelegramChatTypeMap = map[string]TelegramChatType{
	"channel": Channel,
	"group":   Group,
}

func ParseTelegramChatType(str string) (TelegramChatType, bool) {
	c, ok := TelegramChatTypeMap[strings.ToLower(str)]
	return c, ok
}

type TelegramClient struct {
	bot        *tgbotapi.BotAPI
	chatType   TelegramChatType
	credential string
}

func (client TelegramClient) SendText(text string) error {
	messageConfig := tgbotapi.MessageConfig{}

	if client.chatType == Channel {
		messageConfig = tgbotapi.NewMessageToChannel(client.credential, text)
	}

	if client.chatType == Group {
		chatID, err := strconv.Atoi(client.credential)
		if err != nil {
			return err
		}

		messageConfig = tgbotapi.NewMessage(int64(chatID), text)
	}

	messageConfig.ParseMode = "markdown"
	messageConfig.DisableWebPagePreview = true

	_, err := client.bot.Send(messageConfig)
	if err != nil {
		return err
	}

	return nil
}

func (client TelegramClient) SendImage(imageUrl string) error {
	imageConfig := tgbotapi.PhotoConfig{}

	if client.chatType == Channel {
		imageConfig = tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat:    tgbotapi.BaseChat{ChannelUsername: client.credential},
				FileID:      imageUrl,
				UseExisting: true,
			},
		}
	}

	if client.chatType == Group {
		chatID, err := strconv.Atoi(client.credential)
		if err != nil {
			return err
		}

		imageConfig = tgbotapi.NewPhotoShare(int64(chatID), imageUrl)
	}

	_, err := client.bot.Send(imageConfig)
	if err != nil {
		return err
	}

	return nil
}

func NewTelegramClient(token, credential string, chatType TelegramChatType) (Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return TelegramClient{}, err
	}

	return TelegramClient{
		bot:        bot,
		chatType:   chatType,
		credential: credential,
	}, nil
}
