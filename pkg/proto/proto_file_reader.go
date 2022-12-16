package proto

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var inMessage bool
var inEnum bool

type ProtobufFileReader struct {
	pkg           *Package
	inMessage     bool
	inEnum        bool
	commentBuffer strings.Builder
	messageBuffer []string
	enumBuffer    []string
}

func NewReadState() *ProtobufFileReader {
	return &ProtobufFileReader{
		inMessage:     false,
		inEnum:        false,
		commentBuffer: strings.Builder{},
		messageBuffer: make([]string, 0),
		enumBuffer:    make([]string, 0),
	}
}

func (rs *ProtobufFileReader) AppendMessageLine(in string) {
	rs.messageBuffer = append(rs.messageBuffer, in)
}

func (rs *ProtobufFileReader) AppendEnumLine(in string) {
	rs.enumBuffer = append(rs.enumBuffer, in)
}

func (rs *ProtobufFileReader) ResetComments() string {
	out := rs.commentBuffer.String()
	rs.commentBuffer.Reset()
	return out
}

func (rs *ProtobufFileReader) ResetMessageBuffer() []string {
	out := make([]string, len(rs.messageBuffer))
	copy(out, rs.messageBuffer)
	rs.messageBuffer = nil
	rs.messageBuffer = make([]string, 0)
	return out
}

func (rs *ProtobufFileReader) ResetEnumBuffer() []string {
	out := make([]string, len(rs.enumBuffer))
	copy(out, rs.enumBuffer)
	rs.enumBuffer = nil
	rs.enumBuffer = make([]string, 0)
	return out
}

func (rs *ProtobufFileReader) Reset() {
	rs.inMessage = false
	rs.inEnum = false
	rs.commentBuffer.Reset()
}

func (rs *ProtobufFileReader) InitPackage(in string) {
	rs.pkg = NewPackage(in, rs.ResetComments())
}

func (rs *ProtobufFileReader) Read(fileName string) *Package {
	fil, err := os.Open(fileName)
	if err != nil {
		log.Printf("Failed to read file: %s\n", fileName)
	}

	scanner := bufio.NewScanner(fil)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "package") {
			rs.InitPackage(line)
		} else if strings.HasPrefix(line, CommentSingleLine) {
			// Single Line Comment
			rs.commentBuffer.WriteString(CleanComment(line))
		} else if strings.HasPrefix(line, CommentOpenMultiline) {
			// Opens a multi-line comment
			rs.commentBuffer.WriteString(CleanComment(line))
		} else if strings.HasPrefix(line, "}") {
			if rs.inMessage {
				msg := NewMessage(rs.pkg.Name, rs.ResetMessageBuffer(), rs.ResetComments())
				if msg != nil {
					rs.pkg.AddMessage(msg)
				}
			} else if rs.inEnum {
				e := NewEnum(rs.pkg.Name, rs.ResetEnumBuffer(), rs.ResetComments())
				if e != nil {
					rs.pkg.AddEnum(e)
				}
			}
			rs.Reset()
		} else if strings.HasPrefix(line, "message ") || rs.inMessage {
			rs.inMessage = true
			rs.AppendMessageLine(line)
		} else if strings.HasPrefix(line, "enum ") || rs.inEnum {
			rs.inEnum = true
			rs.AppendEnumLine(line)
		}
	}

	return rs.pkg
}

const (
	CommentSingleLine     = `//(.*)`
	CommentOpenMultiline  = `/\*(.*)`
	CommentCloseMultiLine = `\*/`
)

func CleanComment(in string) string {
	if strings.HasPrefix(in, "//") {
		return strings.Trim(in[3:], Space)
	} else if strings.HasPrefix(in, "*") {
		return strings.Trim(in[1:], Space)
	} else if strings.HasSuffix(in, CommentCloseMultiLine) {
		return strings.Trim(in[0:len(in)-2], Space)
	}
	return ""
}
