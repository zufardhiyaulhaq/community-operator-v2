package handler

import (
	"github.com/zufardhiyaulhaq/community-operator/pkg/message"
)

type Handler interface {
	SendMessage(message message.Message) error
}
