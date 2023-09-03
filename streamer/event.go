package streamer

type Event struct {
	topic         string // topic name
	sourceAddress string // source of the topic, for nsq its a nsqlookupd address
}

// NewEvent create new event
func NewEvent(topic string, source string) Event {
	return Event{topic: topic, sourceAddress: source}
}

// GetTopic return topic name
func (e Event) GetTopic() string {
	return e.topic
}

// GetSource return the source address for the topic
func (e Event) GetSource() string {
	return e.sourceAddress
}
