package proto

import (
	"bufio"
	"errors"
)

var InvalidImport = errors.New("invalid import")

const (
	Empty         = ""
	Space         = " "
	Equal         = "="
	OpenBracket   = "["
	ClosedBracket = "]"
	DoubleQuote   = `"`
	Semicolon     = ";"
)

type NamedValue struct {
	Name    string
	Value   any
	Comment *Comment
}

type Annotation struct {
	*NamedValue
}

type Qualified struct {
	Qualifier string
	Name      string
	Comment   *Comment
}

var RegisteredVisitors []Visitor

func init() {
	// Handle Package
	RegisteredVisitors = append(RegisteredVisitors, &PackageVisitor{})

	// Handle Comments
	RegisteredVisitors = append(RegisteredVisitors, &CommentVisitor{})

	// Handle Imports
	RegisteredVisitors = append(RegisteredVisitors, &ImportVisitor{})
	// Handle Options
	RegisteredVisitors = append(RegisteredVisitors, &OptionVisitor{})

	// Handle Messages
	RegisteredVisitors = append(RegisteredVisitors, &MessageVisitor{})

	// Must be last as it's the most forgiving
	RegisteredVisitors = append(RegisteredVisitors, &AttributeVisitor{})
}

type Visitor interface {
	CanVisit(in string) bool
	Visit(in string, scanner *bufio.Scanner, comment *Comment) interface{}
}
