package proto

import (
	"fmt"
	"strings"
)

type ImportVisitor struct {
}

func (iv *ImportVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "import ") && in.Token == Semicolon
}

func (iv *ImportVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	fmt.Println("Visiting Import")
	fValues := in.SplitSyntax()
	return NewImport(RemoveDoubleQuotes(RemoveSemicolon(fValues[1])))
}
