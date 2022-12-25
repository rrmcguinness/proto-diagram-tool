package proto

import (
	"strings"
)

// CommentVisitor is responsible for reading and recoding comment lines,
// it DOES NOT handle inline comments because of how the
type CommentVisitor struct {
}

func (cv *CommentVisitor) CanVisit(in *Line) bool {
	return in.Token == InlineCommentPrefix || in.Token == MultiLineCommentInitiator
}

func (cv *CommentVisitor) Visit(scanner Scanner, in *Line, _ string) interface{} {
	Log.Debug("Visiting Comment")
	var out Comment
	if in.Token == MultiLineCommentInitiator {
		// Append subsequent lines until the terminator is reached
		for scanner.Scan() {
			var n = scanner.Text()
			if strings.HasSuffix(n, MultilineCommentTerminator) {
				break
			}
			out = Comment(Join(Space, string(out), n))
		}
	} else if in.Token == InlineCommentPrefix {
		out += Comment(in.Comment)
	}
	return out
}
