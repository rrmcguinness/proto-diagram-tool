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
	v, e := strconv.ParseInt(RemoveSemicolon(ordinal), 10, 64)
	if e != nil {
		fmt.Sprintf("===========> %v", e)
	}
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
	Values []*EnumValue
}

func NewEnum(q string, name string, comment *Comment) *Enum {
	if comment == nil {
		comment = &Comment{}
	}
	return &Enum{
		Qualified: &Qualified{
			Qualifier: q,
			Name:      name,
			Comment:   comment,
		},
		Values: make([]*EnumValue, 0),
	}
}

func (e *Enum) AddEnumValue(ordinal string, name string) {
	e.Values = append(e.Values, NewEnumValue(ordinal, name))
}

func (e *Enum) AddEnumValueWithComment(ordinal string, name string, comment string) {
	e.Values = append(e.Values, NewEnumValueWithComment(ordinal, name, &Comment{Value: comment}))
}

type EnumVisitor struct {
	// Handle Comments
	commentVisitor *CommentVisitor
}

func (ev *EnumVisitor) init() {
	ev.commentVisitor = &CommentVisitor{}
}

func (ev *EnumVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "enum ") && strings.HasSuffix(in, OpenBrace)
}

func (ev *EnumVisitor) Visit(namespace string, in string, scanner *bufio.Scanner, comment *Comment) interface{} {
	fmt.Println("Visiting Enum")
	in = CleanSpaces(in)
	fValues := strings.Split(in, Space)

	out := NewEnum(namespace, fValues[1], comment)

	for scanner.Scan() {
		n := CleanSpaces(scanner.Text())
		if strings.HasSuffix(n, CloseBrace) {
			break
		}

		eValues := strings.Split(n, Space)
		if len(eValues) == 3 {
			out.AddEnumValue(eValues[2], eValues[0])
		} else if len(eValues) > 3 {
			cmt := strings.Join(eValues[4:], Space)
			out.AddEnumValueWithComment(eValues[2], eValues[0], cmt)
		}
	}
	return out
}
