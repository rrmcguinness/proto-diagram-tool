package proto

import (
	"strings"
)

// NewAttributeVisitor - Constructor for the AttributeVisitor
func NewAttributeVisitor() *attributeVisitor {
	Log.Debug("Initializing attributeVisitor")
	return &attributeVisitor{}
}

// Visitor implementation for attributes
type attributeVisitor struct {
}

// CanVisit - Determines if the line is an attribute, it doesn't end in a brace,
// it's a map, repeated, or can effectively be split
func (av attributeVisitor) CanVisit(in *Line) bool {
	return (!strings.HasSuffix(in.Syntax, OpenBrace) || !strings.HasSuffix(in.Syntax, CloseBrace)) &&
		strings.HasPrefix(in.Syntax, "repeated") ||
		strings.HasPrefix(in.Syntax, "map") || len(in.SplitSyntax()) >= 4
}

func handleRepeated(out *Attribute, split []string) {
	Log.Debugf("\t processing repeated attribute %s", split[2])
	// 0 - 4 repeated, type, name, equals, ordinal
	out.Repeated = true
	out.Kind = append(out.Kind, split[1])
	out.Name = split[2]
	out.Ordinal = ParseOrdinal(split[4])
}

func handleMap(out *Attribute, split []string) {
	Log.Debugf("\t processing map attribute %s", split[2])
	// map1, map2, name, equals, ordinal
	mapValue := Join(Space, split[0], split[1])
	innerTypes := mapValue[strings.Index(mapValue, OpenMap)+len(OpenMap) : strings.Index(mapValue, CloseMap)]
	splitTypes := strings.Split(innerTypes, Comma)
	out.Name = split[2]
	out.Map = true
	out.Kind = append(out.Kind, splitTypes...)
	out.Ordinal = ParseOrdinal(split[4])
}

func handleDefaultAttribute(out *Attribute, split []string) {
	Log.Debugf("\t processing standard attribute %s", split[1])
	out.Name = split[1]
	out.Kind = append(out.Kind, split[0])
	out.Ordinal = ParseOrdinal(split[3])
}

func (av attributeVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	Log.Debug("Visiting Attribute")
	out := NewAttribute(namespace, in.Comment)
	out.Annotations = ParseAnnotations(in.Syntax)
	split := in.SplitSyntax()

	if strings.HasPrefix(in.Syntax, Reserved) {
		Log.Debug("\t processing reserved attribute")
		out.Comment += Space + in.Comment
	} else if strings.HasPrefix(in.Syntax, Repeated) {
		handleRepeated(out, split)
	} else if strings.HasPrefix(in.Syntax, Map) {
		handleMap(out, split)
	} else {
		handleDefaultAttribute(out, split)
	}
	return out
}

func ParseAnnotations(in string) []*Annotation {
	Log.Debug("Processing Annotation")
	out := make([]*Annotation, 0)
	if strings.Contains(in, OpenBracket) && strings.Contains(in, ClosedBracket) {
		annotationString := in[strings.Index(in, OpenBracket)+1 : strings.Index(in, ClosedBracket)]
		split := strings.Split(strings.ReplaceAll(annotationString, SingleQuote, Empty), Space)
		out = append(out, &Annotation{
			Name:  split[0],
			Value: split[2],
		})
	}
	return out
}
