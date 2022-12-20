package proto

import "strings"

func RemoveSemicolon(in string) string {
	return strings.ReplaceAll(in, ";", Empty)
}
