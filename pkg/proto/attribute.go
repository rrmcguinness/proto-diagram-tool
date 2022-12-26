package proto

// An Attribute is a component in the message structure.
type Attribute struct {
	*Qualified
	Repeated    bool
	Map         bool
	Kind        []string
	Ordinal     int
	Annotations []*Annotation
}

// IsValid implements the Validatable interface
func (a *Attribute) IsValid() bool {
	return len(a.Name) > 0 && a.Kind != nil && len(a.Kind) >= 1 && a.Ordinal >= 1
}

// NewAttribute is the Attribute constructor
func NewAttribute(namespace string, comment Comment) *Attribute {
	return &Attribute{
		Qualified:   &Qualified{Qualifier: namespace, Comment: comment},
		Repeated:    false,
		Annotations: make([]*Annotation, 0)}
}
