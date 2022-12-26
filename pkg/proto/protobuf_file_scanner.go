package proto

import (
	"bufio"
	"os"
	"regexp"
)

var SpaceRemover *regexp.Regexp

func init() {
	SpaceRemover = regexp.MustCompile(SpaceRemovalRegex)
}

// NewProtobufFileScanner is the constructor for ProtobufFileScanner
func NewProtobufFileScanner(file *os.File) Scanner {
	return &ProtobufFileScanner{scanner: bufio.NewScanner(file)}
}

// ProtobufFileScanner is a specialized scanner for reading protobuf 3 files.
type ProtobufFileScanner struct {
	scanner *bufio.Scanner
}

// Scan is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Scan() bool {
	return sw.scanner.Scan()
}

// Text is a specialization of the Text function, ensuring the line read
// is ready for processing.
func (sw ProtobufFileScanner) Text() string {
	return FormatLine(sw.scanner.Text())
}

// Split is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Split(splitFunction bufio.SplitFunc) {
	sw.scanner.Split(splitFunction)
}

// Buffer is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Buffer(buf []byte, max int) {
	sw.scanner.Buffer(buf, max)
}

// Err is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Err() error {
	return sw.scanner.Err()
}

// Bytes is a delegate method to the underline scanner
func (sw ProtobufFileScanner) Bytes() []byte {
	return sw.scanner.Bytes()
}

// ReadLine is an addition to the buffered reader responsible for interpreting
// the line of the protobuf for the AST.
func (sw ProtobufFileScanner) ReadLine() *Line {
	return NewLine(sw.Text())
}
