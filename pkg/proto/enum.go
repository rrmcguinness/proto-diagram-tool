package proto

// Enum represents a Proto Enum type.
type Enum struct {
	*Qualified
	Values []*EnumValue
}

// NewEnum is the Enum Constructor
func NewEnum(q string, name string, comment Comment) *Enum {
	return &Enum{
		Qualified: &Qualified{
			Qualifier: q,
			Name:      name,
			Comment:   comment,
		},
		Values: make([]*EnumValue, 0),
	}
}
