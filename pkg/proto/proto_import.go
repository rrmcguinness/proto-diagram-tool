package proto

import (
	"bufio"
	"fmt"
	"strings"
)

type Import struct {
	Path string
}

func NewImport(path string) *Import {
	return &Import{Path: path}
}

type ImportVisitor struct {
}

func (iv *ImportVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "import ") && strings.HasSuffix(in, Semicolon)
}

func (iv *ImportVisitor) Visit(_ string, in string, _ *bufio.Scanner, comment *Comment) interface{} {
	fmt.Println("Visiting Import")
	fValues := strings.Split(in, Space)
	return NewImport(RemoveDoubleQuotes(RemoveSemicolon(fValues[1])))
}
