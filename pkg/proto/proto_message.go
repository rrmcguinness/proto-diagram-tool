package proto

import (
	"bufio"
	"fmt"
	"strings"
)

type Message struct {
	*Qualified
	Attributes []*Attribute
	Messages   []*Message
	Enums      []*Enum
}

func NewMessage() *Message {
	return &Message{
		Attributes: make([]*Attribute, 0),
		Messages:   make([]*Message, 0),
		Enums:      make([]*Enum, 0),
	}
}

type MessageVisitor struct {
}

func (mv *MessageVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "message ")
}

func (mv *MessageVisitor) Visit(in string, scanner *bufio.Scanner, comment *Comment) interface{} {
	fmt.Printf("Visiting Message: %s", in)

	arr := strings.Split(in, Space)
	out := &Message{Qualified: &Qualified{Name: arr[1], Comment: comment}}
	if n := scanner.Text(); strings.HasSuffix(n, "}") {
		for _, visitor := range RegisteredVisitors {
			var elementComment *Comment

			if visitor.CanVisit(n) {
				rt := visitor.Visit(n, scanner, elementComment)
				switch t := rt.(type) {
				case Message:
					out.Messages = append(out.Messages, &t)
				case Enum:
					out.Enums = append(out.Enums, &t)
				case Attribute:
					out.Attributes = append(out.Attributes, &t)
				case Comment:
					elementComment = &t
				}
			}
		}
	}
	return out
}
