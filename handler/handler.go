package handler

import (
	"github.com/nyelonong/nsqsink/message"
)

type Handler interface {
	Handle(msg message.Message) error
}
