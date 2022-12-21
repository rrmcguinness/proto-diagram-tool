package proto

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var SpaceRemover *regexp.Regexp

func init() {
	SpaceRemover = regexp.MustCompile(SpaceRemovalRegex)
}

func RemoveSemicolon(in string) string {
	return strings.ReplaceAll(in, Semicolon, Empty)
}

func RemoveDoubleQuotes(in string) string {
	return strings.ReplaceAll(in, DoubleQuote, Empty)
}

func StripComment(in string) string {
	if strings.Contains(in, InlineCommentPrefix) {
		return in[strings.Index(in, InlineCommentPrefix)+len(InlineCommentPrefix):]
	}
	return ""
}

func ParseOrdinal(in string) int {
	i, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		fmt.Printf("Failed to parse %s for integer", in)
		return 0
	}
	return int(i)
}

func CleanSpaces(in string) string {
	return strings.TrimSpace(SpaceRemover.ReplaceAllString(in, " "))
}

func Join(joinCharacter string, values ...string) string {
	out := ""
	count := len(values)
	for i := 0; i < count; i++ {
		if i < count-1 {
			out += values[i] + joinCharacter
		} else {
			out += values[i]
		}
	}
	return out
}
