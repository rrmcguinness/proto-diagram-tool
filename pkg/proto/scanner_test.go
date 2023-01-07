package proto

import (
	"bufio"
	"strings"
)

type TestScanner struct {
	internalScanner *bufio.Scanner
}

func (ts *TestScanner) Scan() bool {
	return ts.internalScanner.Scan()
}

func (ts *TestScanner) Text() string {
	return ts.internalScanner.Text()
}

func (ts *TestScanner) Split(splitFunction bufio.SplitFunc) {
	ts.internalScanner.Split(splitFunction)
}

func (ts *TestScanner) Buffer(buf []byte, max int) {
	ts.internalScanner.Buffer(buf, max)
}

func (ts *TestScanner) Err() error {
	return ts.Err()
}

func (ts *TestScanner) Bytes() []byte {
	return ts.Bytes()
}

func (ts *TestScanner) ReadLine() *Line {
	return NewLine(ts.internalScanner.Text())
}

func NewTestScanner(in string) *TestScanner {
	return &TestScanner{
		internalScanner: bufio.NewScanner(strings.NewReader(in)),
	}
}
