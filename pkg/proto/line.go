package proto

import "strings"

// Line is a split line for syntax, token, and comment
type Line struct {
	Syntax  string
	Token   string
	Comment Comment
}

func NewLine(in string) *Line {
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

// SplitSyntax breaks the syntax line on Space the character
func (l *Line) SplitSyntax() []string {
	return strings.Split(l.Syntax, Space)
}
