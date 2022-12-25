package proto

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var SpaceRemover *regexp.Regexp

func init() {
	SpaceRemover = regexp.MustCompile(SpaceRemovalRegex)
}

func FormatLine(in string) string {
	return strings.TrimSpace(SpaceRemover.ReplaceAllString(in, " "))
}

func NewScannerWrapper(file *os.File) Scanner {
	return &ScannerWrapper{scanner: bufio.NewScanner(file)}
}

type ScannerWrapper struct {
	scanner *bufio.Scanner
}

func (sw ScannerWrapper) Scan() bool {
	return sw.scanner.Scan()
}

func (sw ScannerWrapper) Text() string {
	return FormatLine(sw.scanner.Text())
}

func (sw ScannerWrapper) Split(splitFunction bufio.SplitFunc) {
	sw.scanner.Split(splitFunction)
}

func (sw ScannerWrapper) Buffer(buf []byte, max int) {
	sw.scanner.Buffer(buf, max)
}

func (sw ScannerWrapper) Err() error {
	return sw.scanner.Err()
}

func (sw ScannerWrapper) Bytes() []byte {
	return sw.scanner.Bytes()
}

func (sw ScannerWrapper) ReadLine() *Line {
	in := sw.Text()
	line := &Line{}
	if strings.HasPrefix(in, InlineCommentPrefix) {
		// Handle single comments
		line.Comment = Comment(strings.TrimSpace(in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):]))
		line.Token = InlineCommentPrefix
	} else if strings.HasPrefix(in, MultiLineCommentInitiator) {
		// Handle Multiline Comments
		line.Comment = Comment(strings.TrimSpace(in[strings.Index(in, MultiLineCommentInitiator)+len(MultiLineCommentInitiator):]))
		line.Token = MultiLineCommentInitiator
	} else if strings.Contains(in, Semicolon) {
		// Handle Syntax Stings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, Semicolon)])
		line.Token = Semicolon
	} else if strings.Contains(in, OpenBrace) {
		// Handle Structure Strings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, OpenBrace)])
		line.Token = OpenBrace
	} else if strings.Contains(in, CloseBrace) {
		// Handle Inline Closed Structure Strings
		line.Syntax = strings.TrimSpace(in[0:strings.Index(in, CloseBrace)])
		line.Token = CloseBrace
	}
	// Add Inline Comments
	if !strings.HasPrefix(in, InlineCommentPrefix) && strings.Contains(in, InlineCommentPrefix) {
		line.Comment += Comment(Space + strings.TrimSpace(in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):]))
	}
	line.Comment = line.Comment.TrimSpace()
	return line
}
