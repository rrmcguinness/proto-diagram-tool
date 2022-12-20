package proto

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	InlineCommentPrefix        = "//"
	MultiLineCommentInitiator  = "/*"
	MultilineCommentTerminator = "*/"
	PrefixedComment            = "* "
)

type Comment struct {
	value string
}

type CommentVisitor struct {
}

func (cv *CommentVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, InlineCommentPrefix) || strings.HasPrefix(in, MultiLineCommentInitiator)
}

func (cv *CommentVisitor) Visit(in string, scanner *bufio.Scanner, _ *Comment) interface{} {
	fmt.Println("Visiting Comment")

	out := &Comment{value: strings.TrimSpace(in[len(InlineCommentPrefix):])}

	if strings.HasPrefix(in, MultiLineCommentInitiator) {
		// Append subsequent lines until the terminator is reached
		for scanner.Scan() {
			var n = strings.TrimSpace(scanner.Text())
			fmt.Println(n)
			if strings.HasSuffix(n, MultilineCommentTerminator) {
				break
			}
			if strings.HasPrefix(n, PrefixedComment) {
				out.value += strings.TrimSpace(n[len(PrefixedComment):])
			} else {
				out.value += strings.TrimSpace(n)
			}
		}
	}

	fmt.Printf("Comment: %s\n", out.value)
	return out
}
