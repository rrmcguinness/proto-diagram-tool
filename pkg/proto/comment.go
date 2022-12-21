package proto

import (
	"bufio"
	"fmt"
	"strings"
)

type Comment struct {
	Value string
}

type CommentVisitor struct {
}

func (cv *CommentVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, InlineCommentPrefix) || strings.HasPrefix(in, MultiLineCommentInitiator)
}

func (cv *CommentVisitor) Visit(_ string, in string, scanner *bufio.Scanner, _ *Comment) interface{} {
	fmt.Println("Visiting Comment")

	out := &Comment{Value: strings.TrimSpace(in[len(InlineCommentPrefix):])}

	if strings.HasPrefix(in, MultiLineCommentInitiator) {
		// Append subsequent lines until the terminator is reached
		for scanner.Scan() {
			var n = strings.TrimSpace(scanner.Text())
			if strings.HasSuffix(n, MultilineCommentTerminator) {
				break
			}
			if strings.HasPrefix(n, PrefixedComment) {
				out.Value += strings.TrimSpace(n[len(PrefixedComment):])
			} else {
				out.Value += strings.TrimSpace(n)
			}
		}
	}
	return out
}
