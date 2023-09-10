package message

import "time"

type Message interface {
	Finish()
	Requeue(delay time.Duration)
	RequeueWithoutBackoff(delay time.Duration)
	Attempts() uint16
	Body() []byte
}
