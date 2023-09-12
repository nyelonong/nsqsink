package message

import (
	"time"

	"github.com/nsqio/go-nsq"
)

// Messager
// contract for message data
type Messager interface {
	Finish()
	Requeue(delay time.Duration)
	RequeueWithoutBackoff(delay time.Duration)
	GetAttempts() uint16
	GetBody() []byte
}

type Message struct {
	*nsq.Message
}

// New
// return message object
func New(msg *nsq.Message) Messager {
	return Message{msg}
}
