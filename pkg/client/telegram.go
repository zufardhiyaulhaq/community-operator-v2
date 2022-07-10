package client

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramChatType int64

const (
	Channel TelegramChatType = 0
	Group   TelegramChatType = 1
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

func (client TelegramClient) Send(c Chattable) (Response, error) {
	if c.Method() == "sendText" {
		messageConfig := tgbotapi.MessageConfig{}

		if client.chatType == Channel {
			messageConfig = tgbotapi.NewMessageToChannel(client.credential, c.Value())
		}

		if client.chatType == Group {
			chatID, err := strconv.Atoi(client.credential)
			if err != nil {
				return Response{}, err
			}

			messageConfig = tgbotapi.NewMessage(int64(chatID), c.Value())
		}

		messageConfig.ParseMode = "markdown"
		messageConfig.DisableWebPagePreview = true

		response, err := client.bot.Send(messageConfig)
		if err != nil {
			return Response{}, err
		}

		return Response{
			Identifier: string(response.MessageID),
		}, nil
	}

	if c.Method() == "sendImageUrl" {
		imageConfig := tgbotapi.PhotoConfig{}

		if client.chatType == Channel {
			imageConfig = tgbotapi.PhotoConfig{
				BaseFile: tgbotapi.BaseFile{
					BaseChat:    tgbotapi.BaseChat{ChannelUsername: client.credential},
					FileID:      c.Value(),
					UseExisting: true,
				},
			}
		}

		if client.chatType == Group {
			chatID, err := strconv.Atoi(client.credential)
			if err != nil {
				return Response{}, err
			}

			imageConfig = tgbotapi.NewPhotoShare(int64(chatID), c.Value())
		}

		response, err := client.bot.Send(imageConfig)
		if err != nil {
			return Response{}, err
		}

		return Response{
			Identifier: string(response.MessageID),
		}, nil
	}

	return Response{}, fmt.Errorf("chattable method not found")
}

func (client TelegramClient) Reply(c Chattable, identifier string) (Response, error) {
	return Response{}, fmt.Errorf("chattable method not found")
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
