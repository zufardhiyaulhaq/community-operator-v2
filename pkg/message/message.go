package message

import (
	"bytes"
)

type HandlerType int64

const (
	Telegram HandlerType = 0
	Twitter  HandlerType = 1
)

type Message interface {
	RenderText(handler HandlerType) (bytes.Buffer, error)
	RenderImageUrl() (string, error)
	SendText() bool
	SendImage() bool
}
