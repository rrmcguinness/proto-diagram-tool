package proto

import "bufio"

// Validatable is a reference interface for the validator pattern
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
