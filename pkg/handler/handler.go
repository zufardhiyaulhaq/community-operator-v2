package handler

import (
	"github.com/zufardhiyaulhaq/community-operator-v2/pkg/message"
)

type Handler interface {
	SendMessage(message message.Message) error
}
