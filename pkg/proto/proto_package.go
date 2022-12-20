package proto

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var isDebug bool

type Package struct {
	Path     string
	Name     string
	Comment  *Comment
	Options  []*Option
	Imports  []*Import
	Messages []*Message
	Enums    []*Enum
}

func NewPackage(path string) *Package {
	pkg := &Package{Path: path}
	return pkg
}

func (p *Package) Read(debug bool) error {
	isDebug = debug

	readFile, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		fmt.Printf("Current Line: `%s`\n", line)

		for _, visitor := range RegisteredVisitors {
			var comment *Comment
			if visitor.CanVisit(line) {
				rt := visitor.Visit(line, scanner, comment)
				switch t := rt.(type) {
				case *Option:
					p.Options = append(p.Options, t)
				case *Import:
					p.Imports = append(p.Imports, t)
				case *Message:
					p.Messages = append(p.Messages, t)
				case *Enum:
					p.Enums = append(p.Enums, t)
				case *Comment:
					comment = t
				case *Package:
					p.Name = t.Name
					if comment != nil {
						p.Comment = &Comment{value: comment.value}
					}
				default:
					fmt.Printf("Unhandled Return type for package: %T visitor\n", t)
				}
			}
		}
	}

	return nil
}

type PackageVisitor struct {
}

func (pv *PackageVisitor) CanVisit(in string) bool {
	return strings.HasPrefix(in, "package ") && strings.HasSuffix(in, Semicolon)
}

func (pv *PackageVisitor) Visit(in string, _ *bufio.Scanner, _ *Comment) interface{} {
	fValues := strings.Split(in, Space)
	return &Package{Name: RemoveSemicolon(fValues[1])}
}
