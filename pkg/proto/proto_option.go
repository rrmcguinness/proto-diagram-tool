package proto

import (
	"bufio"
	"fmt"
	"strings"
)

type Option struct {
	*NamedValue
}

type OptionVisitor struct {
}

func (ov *OptionVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "option ") && strings.HasSuffix(in, Semicolon)
}

func (ov *OptionVisitor) Visit(in string, _ *bufio.Scanner, comment *Comment) interface{} {
	fmt.Println("Visiting Option")

	fValues := strings.Split(in, Space)
	if len(fValues) == 4 {
		return &Option{&NamedValue{
			Name:    fValues[1],
			Value:   fValues[3],
			Comment: comment,
		}}
	}
	return &Option{
		NamedValue: &NamedValue{Name: "Invalid"},
	}
}
