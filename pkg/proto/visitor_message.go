package proto

import (
	"strings"
)

// MessageVisitor is used for interpreting message text
type MessageVisitor struct {
}

// CanVisit visits if the line starts with 'message' and ends with an open brace '{'
func (mv *MessageVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "message ") && in.Token == OpenBrace
}

// Visit evaluates the current line and parses the message until the closed brace
// is evaluated.
func (mv *MessageVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting Message: %s\n", in)

	values := in.SplitSyntax()
	out := NewMessage()

	out.Name = values[1]
	out.Comment = in.Comment

	var comment = Comment("")

	for scanner.Scan() {
		line := scanner.ReadLine()

		Log.Debugf("Current Line: `%s`\n", line)

		if strings.HasSuffix(line.Token, CloseBrace) {
			break
		}
		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(line) {
				rt := visitor.Visit(
					scanner,
					line,
					Join(Period, namespace, out.Name))
				switch t := rt.(type) {
				case *Message:
					t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
					out.Messages = append(out.Messages, t)
					comment = comment.Clear()
				case *Enum:
					t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
					out.Enums = append(out.Enums, t)
					comment = comment.Clear()
				case *Attribute:
					if t.IsValid() {
						t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
						out.Attributes = append(out.Attributes, t)
						comment = comment.Clear()
					}
				case Comment:
					comment = comment.Append(t).AddSpace()
				}
			}
		}
	}
	return out
}
