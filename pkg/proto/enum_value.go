package proto

// EnumValue is the representation of an internal enum value.
type EnumValue struct {
	Namespace string
	Ordinal   int
	Value     string
	Comment   Comment
}

// NewEnumValue is the EnumValue constructor
func NewEnumValue(namespace string, ordinal string, value string, comment Comment) *EnumValue {
	return &EnumValue{Namespace: namespace, Ordinal: ParseOrdinal(ordinal), Value: value, Comment: comment}
}
