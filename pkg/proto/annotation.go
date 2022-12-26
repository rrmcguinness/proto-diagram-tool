package proto

import "strings"

// Annotation is an inline structure applicable only to attributes
type Annotation struct {
	Name  string
	Value any
}

// NewAnnotation is the Annotation Constructor
func NewAnnotation(name string, value any) *Annotation {
	return &Annotation{Name: name, Value: value}
}

func ParseAnnotations(in string) []*Annotation {
	Log.Debug("Processing Annotation")
	out := make([]*Annotation, 0)
	if strings.Contains(in, OpenBracket) && strings.Contains(in, ClosedBracket) {
		annotationString := in[strings.Index(in, OpenBracket)+1 : strings.Index(in, ClosedBracket)]
		split := strings.Split(strings.ReplaceAll(annotationString, SingleQuote, Empty), Space)
		out = append(out, NewAnnotation(split[0], split[2]))
	}
	return out
}
