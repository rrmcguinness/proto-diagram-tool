package proto

import "strings"

// Comment is a string with additional methods
type Comment string

// Append adds a comment to the end of an existing comment.
func (c Comment) Append(other Comment) Comment {
	c += Comment(strings.TrimSpace(string(other)))
	return c
}

// AddSpace adds a space to the end of a Comment.
func (c Comment) AddSpace() Comment {
	c += Space
	return c
}

// TrimSpace removes any double space or padding spaces from the comment.
func (c Comment) TrimSpace() Comment {
	return Comment(FormatLine(strings.TrimSpace(string(c))))
}

// Clear truncates the comment
func (c Comment) Clear() Comment {
	return c[:0]
}
