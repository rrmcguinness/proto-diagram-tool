package proto

import (
	"fmt"
	"strings"
)

type Option struct {
	*NamedValue
}

type OptionVisitor struct {
}

func (ov *OptionVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "option ") && in.Token == Semicolon
}

func (ov *OptionVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	fmt.Println("Visiting Option")
	fValues := in.SplitSyntax()
	if len(fValues) == 4 {
		return &Option{&NamedValue{
			Name:    fValues[1],
			Value:   fValues[3],
			Comment: in.Comment[:],
		}}
	}
	return &Option{
		NamedValue: &NamedValue{Name: "Invalid"},
	}
}
