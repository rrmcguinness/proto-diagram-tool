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
		Qualified:  &Qualified{Comment: &Comment{}},
		Attributes: make([]*Attribute, 0),
		Messages:   make([]*Message, 0),
		Enums:      make([]*Enum, 0),
	}
}

type MessageVisitor struct {
}

func (mv *MessageVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "message ") && strings.HasSuffix(in, OpenBrace)
}

func (mv *MessageVisitor) Visit(namespace string, in string, scanner *bufio.Scanner, comment *Comment) interface{} {
	fmt.Printf("Visiting Message: %s\n", in)

	arr := strings.Split(in, Space)
	out := NewMessage()
	out.Name = arr[1]
	out.Comment = comment

	var elementComment *Comment

	for scanner.Scan() {
		n := CleanSpaces(scanner.Text())
		if strings.HasSuffix(n, CloseBrace) {
			break
		}
		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(n) {
				rt := visitor.Visit(Join(Period, namespace, out.Name), n, scanner, elementComment)
				switch t := rt.(type) {
				case *Message:
					out.Messages = append(out.Messages, t)
				case *Enum:
					out.Enums = append(out.Enums, t)
				case *Attribute:
					if t.IsValid() {
						out.Attributes = append(out.Attributes, t)
					}
				case *Comment:
					elementComment = t
				}
			}
		}
	}
	return out
}
