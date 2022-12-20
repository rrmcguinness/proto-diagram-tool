package proto

import (
	"bufio"
	"fmt"
	"strings"
)

type Attribute struct {
	*Qualified
	Repeated    bool
	Kind        string
	Ordinal     string
	Annotations []*Annotation
}

func NewAttribute(q string, name string, comment *Comment, repeated bool, kind string, ordinal string) *Attribute {
	return &Attribute{Qualified: &Qualified{
		Qualifier: q,
		Name:      name,
		Comment:   comment,
	},
		Repeated: repeated,
		Kind:     kind,
		Ordinal:  ordinal, Annotations: make([]*Annotation, 0)}
}

type AttributeVisitor struct {
}

func (av AttributeVisitor) CanVisit(in string) bool {
	count := len(strings.Split(in, Space))
	return count >= 4 && count <= 8 && strings.HasSuffix(in, Semicolon)
}

func (av AttributeVisitor) Visit(in string, _ *bufio.Scanner, comment *Comment) interface{} {
	fmt.Println("Visiting Attribute")

	values := strings.Split(in, Space)
	count := len(values)
	switch count {
	case 4:
		// No annotations
		return NewAttribute("", values[1], comment, false, values[0], values[3])
	case 5:
		//repeated
		return NewAttribute("", values[2], comment, true, values[1], values[4])
	case 7:
		// Annotation
		// TODO  - Add annotation parser
		return NewAttribute("", values[1], comment, false, values[0], values[3])
	case 8:
		//Repeated with annotation
		// TODO  - Add annotation parser
		return NewAttribute("", values[2], comment, true, values[1], values[4])
	}
	return nil
}
