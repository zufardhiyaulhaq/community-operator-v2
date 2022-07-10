package handler

import (
	"time"

	"github.com/zufardhiyaulhaq/community-operator-v2/pkg/client"
	"github.com/zufardhiyaulhaq/community-operator-v2/pkg/message"
)

type TelegramHandler struct {
	client      client.Client
	handlerType message.HandlerType
}

func (handler TelegramHandler) SendMessage(message message.Message) error {
	if message.SendText() {
		text, err := message.RenderText(handler.handlerType)
		if err != nil {
			return err
		}

		_, err = handler.client.Send(client.Text{
			Text: text.String(),
		})
		if err != nil {
			return err
		}
	}

	if message.SendImage() {
		time.Sleep(30 * time.Second)

		imageUrl, err := message.RenderImageUrl()
		if err != nil {
			return err
		}

		_, err = handler.client.Send(client.ImageUrl{
			ImageUrl: imageUrl,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func NewTelegramHandler(client client.Client) Handler {
	return TelegramHandler{
		client:      client,
		handlerType: message.Telegram,
	}
}
