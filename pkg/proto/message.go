/*
 * Copyright 2022 Google, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proto

import (
	"fmt"
	"strings"
)

// Message represents a message / struct body
type Message struct {
	*Qualified
	Attributes []*Attribute
	Messages   []*Message
	Enums      []*Enum
	Reserved   []*Reserved
}

// NewMessage creates a new message
func NewMessage() *Message {
	return &Message{
		Qualified:  &Qualified{},
		Attributes: make([]*Attribute, 0),
		Messages:   make([]*Message, 0),
		Enums:      make([]*Enum, 0),
		Reserved:   make([]*Reserved, 0),
	}
}

func (m *Message) HasAttributes() bool {
	return len(m.Attributes) > 0
}

func (m *Message) HasMessages() bool {
	return len(m.Messages) > 0
}

func (m *Message) HasEnums() bool {
	return len(m.Enums) > 0
}

func (m *Message) ToMermaid() string {
	out := fmt.Sprintf("\n%s\nclass %s {\n", m.Comment.ToMermaid(), m.Name)
	for _, a := range m.Attributes {
		out += fmt.Sprintf("  %s\n", a.ToMermaid())
	}
	out += "}\n"

	// Handle Attribute Relationships
	for _, a := range m.Attributes {
		if len(a.Kind) == 1 {
			if !strings.Contains(Proto3Types, a.Kind[0]) {
				out += fmt.Sprintf("%s --> `%s`\n", m.Name, a.Kind[0])
			}
		} else if len(a.Kind) == 2 {
			if !strings.Contains(Proto3Types, strings.TrimSpace(a.Kind[1])) {
				out += fmt.Sprintf("%s .. `%s`\n", m.Name, a.Kind[1])
			}
		}
	}

	// Handle Message Relationships
	if m.HasMessages() {
		for _, msg := range m.Messages {
			out += fmt.Sprintf("%s --o `%s`\n", m.Name, msg.Name)
			out += msg.ToMermaid()
		}
	}

	// Handle Enumeration Relationships
	for _, e := range m.Enums {
		out += fmt.Sprintf("%s --o `%s`\n", m.Name, e.Name)
		out += e.ToMermaid()
	}

	return out
}
