package proto

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type EnumValue struct {
	Ordinal int
	Value   string
	Comment *Comment
}

func NewEnumValue(ordinal string, value string) *EnumValue {
	v, _ := strconv.ParseInt(ordinal, 10, 32)
	return &EnumValue{
		Ordinal: int(v),
		Value:   value,
	}
}

func NewEnumValueWithComment(ordinal string, value string, comment *Comment) *EnumValue {
	ev := NewEnumValue(ordinal, value)
	ev.Comment = comment
	return ev
}

type Enum struct {
	*Qualified
	values []*EnumValue
}

func NewEnum(q string, name string, comment *Comment) *Enum {
	return &Enum{
		Qualified: &Qualified{
			Qualifier: q,
			Name:      name,
			Comment:   comment,
		},
		values: make([]*EnumValue, 0),
	}
}

func (e *Enum) AddEnumValue(ordinal string, name string) {
	e.values = append(e.values, NewEnumValue(ordinal, name))
}

func (e *Enum) AddEnumValueWithComment(ordinal string, name string, comment string) {
	e.values = append(e.values, NewEnumValueWithComment(ordinal, name, &Comment{value: comment}))
}

type EnumVisitor struct {
}

func (ev *EnumVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "enum ")
}

func (ev *EnumVisitor) Visit(in string, scanner *bufio.Scanner, comment *Comment) interface{} {
	fmt.Println("Visiting Enum")

	fValues := strings.Split(in, Space)
	out := NewEnum("", fValues[1], comment)

	for n := scanner.Text(); n != "}"; {
		eValues := strings.Split(n, Space)
		if len(eValues) == 3 {
			out.AddEnumValue(RemoveSemicolon(eValues[2]), eValues[0])
		} else if len(eValues) > 3 {
			cmt := strings.Join(eValues[4:], Space)
			out.AddEnumValueWithComment(RemoveSemicolon(eValues[2]), eValues[0], cmt)
		}
	}
	return out
}
