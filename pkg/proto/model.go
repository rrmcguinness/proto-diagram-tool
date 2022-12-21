package proto

import (
	"bufio"
	"errors"
)

var InvalidImport = errors.New("invalid import")

const (
	Repeated = "repeated"
	Map      = "map"
	Reserved = "reserved"

	SpaceRemovalRegex = `\s+`
	Period            = "."
	Empty             = ""
	Space             = " "
	OpenBrace         = "{"
	CloseBrace        = "}"
	OpenBracket       = "["
	ClosedBracket     = "]"
	Semicolon         = ";"
	Comma             = ","

	InlineCommentPrefix        = "//"
	MultiLineCommentInitiator  = "/*"
	MultilineCommentTerminator = "*/"
	PrefixedComment            = "* "
	OpenMap                    = "map<"
	CloseMap                   = ">"
	DoubleQuote                = "\""
	SingleQuote                = "'"
	EndL                       = "\n"
)

type NamedValue struct {
	Name    string
	Value   any
	Comment *Comment
}

type Annotation struct {
	Name  string
	Value any
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

	// Enums
	RegisteredVisitors = append(RegisteredVisitors, &EnumVisitor{})

	// Must be last as it's the most forgiving
	RegisteredVisitors = append(RegisteredVisitors, &AttributeVisitor{})
}

type Visitor interface {
	CanVisit(in string) bool
	Visit(
		namespace string,
		in string,
		scanner *bufio.Scanner,
		comment *Comment) interface{}
}
