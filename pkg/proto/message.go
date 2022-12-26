package proto

// Message represents a message / struct body
type Message struct {
	*Qualified
	Attributes []*Attribute
	Messages   []*Message
	Enums      []*Enum
}

// NewMessage creates a new message
func NewMessage() *Message {
	return &Message{
		Qualified:  &Qualified{},
		Attributes: make([]*Attribute, 0),
		Messages:   make([]*Message, 0),
		Enums:      make([]*Enum, 0),
	}
}
