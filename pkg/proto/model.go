package proto

import (
	"bufio"
	"strings"
)

// Line is a split line for syntax, token, and comment
type Line struct {
	Syntax  string
	Token   string
	Comment Comment
}

func (l *Line) SplitSyntax() []string {
	return strings.Split(l.Syntax, Space)
}

// Comment is a string with additional methods
type Comment string

func (c Comment) Append(other Comment) Comment {
	c += Comment(strings.TrimSpace(string(other)))
	return c
}

func (c Comment) AddSpace() Comment {
	c += Space
	return c
}

func (c Comment) TrimSpace() Comment {
	return Comment(FormatLine(strings.TrimSpace(string(c))))
}

func (c Comment) Clear() Comment {
	return c[:0]
}

// Package is the top level structure of any protobuf
type Package struct {
	Path     string
	Name     string
	Comment  Comment
	Options  []*Option
	Imports  []*Import
	Messages []*Message
	Enums    []*Enum
}

func NewPackage(path string) *Package {
	pkg := &Package{Path: path,
		Options:  make([]*Option, 0),
		Imports:  make([]*Import, 0),
		Messages: make([]*Message, 0),
		Enums:    make([]*Enum, 0),
	}
	return pkg
}

// Import represents an importable file
type Import struct {
	Path    string
	Comment Comment
}

func NewImport(path string) *Import {
	return &Import{Path: path}
}

// NamedValue is super class to capture names and values for typed lines.
type NamedValue struct {
	Name    string
	Value   any
	Comment Comment
}

// Qualified is a super class to capture namespace aware attributes and enums
type Qualified struct {
	Qualifier string
	Name      string
	Comment   Comment
}

// Annotation is an inline structure applicable only to attributes
type Annotation struct {
	Name  string
	Value any
}

// An Attribute is a component in the message structure.
type Attribute struct {
	*Qualified
	Repeated    bool
	Map         bool
	Kind        []string
	Ordinal     int
	Annotations []*Annotation
}

func (a *Attribute) IsValid() bool {
	return len(a.Name) > 0 && a.Kind != nil && len(a.Kind) >= 1 && a.Ordinal >= 1
}

func NewAttribute(namespace string, comment Comment) *Attribute {
	return &Attribute{
		Qualified:   &Qualified{Qualifier: namespace, Comment: comment},
		Repeated:    false,
		Annotations: make([]*Annotation, 0)}
}

// Message represents a message / struct body
type Message struct {
	*Qualified
	Attributes []*Attribute
	Messages   []*Message
	Enums      []*Enum
}

// NewMessage creates a new message
func NewMessage() *Message {
	return &Message{
		Qualified:  &Qualified{},
		Attributes: make([]*Attribute, 0),
		Messages:   make([]*Message, 0),
		Enums:      make([]*Enum, 0),
	}
}

type Enum struct {
	*Qualified
	Values []*EnumValue
}

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

type EnumValue struct {
	Namespace string
	Ordinal   int
	Value     string
	Comment   Comment
}

func NewEnumValue(namespace string, ordinal string, value string, comment Comment) *EnumValue {
	return &EnumValue{Namespace: namespace, Ordinal: ParseOrdinal(ordinal), Value: value, Comment: comment}
}

// SetDebug is used to enable the debug output, useful for troubleshooting.
func SetDebug(debug bool) {
	Log.debug = debug
}

type Validatable interface {
	IsValid() bool
}

// Visitor is an interface used to determine if a line should be read,
// and if it should be, to read and interpret the line and subsequent lines
// as required.
type Visitor interface {
	CanVisit(in *Line) bool
	Visit(
		scanner Scanner,
		in *Line,
		namespace string) interface{}
}

// Scanner is an interface that SHOULD be a Go interface, but is only an
// implementation. Here, we can use the interface to wrap test cases
// with the same behavior of a bufio.Scanner
type Scanner interface {
	Scan() bool
	Text() string
	Split(splitFunction bufio.SplitFunc)
	Buffer(buf []byte, max int)
	Err() error
	Bytes() []byte
	ReadLine() *Line
}
