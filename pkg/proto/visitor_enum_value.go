package proto

type EnumValueVisitor struct {
}

func (evv EnumValueVisitor) CanVisit(in *Line) bool {
	a := in.SplitSyntax()
	return a != nil && len(a) == 3 && in.Token == Semicolon
}

func (evv EnumValueVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	a := in.SplitSyntax()
	return NewEnumValue(namespace, a[2], a[0], in.Comment)
}
