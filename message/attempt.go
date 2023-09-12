package message

// GetAttempts
// return number of attempts consume this message
func (m Message) GetAttempts() uint16 {
	return m.Attempts
}
