package proto

import (
	"bufio"
	"fmt"
	"strings"
)

type Attribute struct {
	*Qualified
	Repeated    bool
	Map         bool
	Kind        []string
	Ordinal     int
	Annotations []*Annotation
}

func (a *Attribute) IsValid() bool {
	return len(a.Name) > 0 && a.Kind != nil && len(a.Kind) >= 1 && a.Ordinal >= 1
}

type AttributeVisitor struct {
}

func (av AttributeVisitor) CanVisit(in string) bool {
	count := len(strings.Split(in, Space))
	return (!strings.HasSuffix(in, OpenBrace) || !strings.HasSuffix(in, CloseBrace)) && strings.HasPrefix(in, "repeated") ||
		strings.HasPrefix(in, "map") || count >= 4
}

func ParseAnnotations(in string) []*Annotation {
	out := make([]*Annotation, 0)
	if strings.Contains(in, OpenBracket) && strings.Contains(in, ClosedBracket) {
		annotationString := CleanSpaces(in[strings.Index(in, OpenBracket)+1 : strings.Index(in, ClosedBracket)])
		split := strings.Split(strings.ReplaceAll(annotationString, SingleQuote, Empty), Space)
		out = append(out, &Annotation{
			Name:  split[0],
			Value: split[2],
		})
	}
	return out
}

func (av AttributeVisitor) Visit(namespace string, in string, _ *bufio.Scanner, comment *Comment) interface{} {
	fmt.Println("Visiting Attribute")

	in = CleanSpaces(RemoveSemicolon(in))

	if len(in) > 0 && strings.Index(in, Space) > 0 {

		out := &Attribute{
			Qualified:   &Qualified{Qualifier: namespace, Comment: &Comment{}},
			Repeated:    false,
			Annotations: make([]*Annotation, 0),
		}

		out.Annotations = ParseAnnotations(in)
		out.Comment.Value = StripComment(in)
		split := strings.Split(in, Space)

		if strings.HasPrefix(in, Reserved) {
			out.Comment.Value += " " + in
		} else if strings.HasPrefix(in, Repeated) {
			// 0 - 4 repeated, type, name, equals, ordinal
			out.Repeated = true
			out.Kind = append(out.Kind, split[1])
			out.Name = split[2]
			out.Ordinal = ParseOrdinal(split[4])
		} else if strings.HasPrefix(in, Map) {
			// map1, map2, name, equals, ordinal
			mapValue := Join(Space, split[0], split[1])
			innerTypes := mapValue[strings.Index(mapValue, OpenMap)+len(OpenMap) : strings.Index(mapValue, CloseMap)]
			splitTypes := strings.Split(innerTypes, Comma)
			out.Name = split[2]
			out.Map = true
			out.Kind = append(out.Kind, splitTypes...)

			out.Ordinal = ParseOrdinal(split[4])
		} else {
			out.Name = split[1]
			out.Kind = append(out.Kind, split[0])
			out.Ordinal = ParseOrdinal(split[3])
		}
		return out
	}
	return nil
}
