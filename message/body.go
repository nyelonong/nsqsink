package message

// GetBody
// return body of the message in byte
func (m Message) GetBody() []byte {
	return m.Body
}
