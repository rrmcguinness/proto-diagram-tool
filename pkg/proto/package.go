package proto

import (
	"os"
)

// Package is the top level structure of any protobuf
type Package struct {
	Path     string
	Name     string
	Comment  Comment
	Options  []*Option
	Imports  []*Import
	Messages []*Message
	Enums    []*Enum
}

func NewPackage(path string) *Package {
	pkg := &Package{Path: path,
		Options:  make([]*Option, 0),
		Imports:  make([]*Import, 0),
		Messages: make([]*Message, 0),
		Enums:    make([]*Enum, 0),
	}
	return pkg
}

func (p *Package) Read(debug bool) error {
	isDebug = debug

	readFile, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	scanner := NewProtobufFileScanner(readFile)

	var comment = Comment("")

	for scanner.Scan() {
		line := scanner.ReadLine()

		Log.Debugf("Current Line: `%s`\n", line)

		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(line) {
				rt := visitor.Visit(scanner, line, p.Name)
				switch t := rt.(type) {
				case *Option:
					t.Comment = comment.AddSpace().Append(line.Comment).TrimSpace()
					p.Options = append(p.Options, t)
					comment = comment.Clear()
				case *Import:
					t.Comment = comment.AddSpace().Append(line.Comment).TrimSpace()
					p.Imports = append(p.Imports, t)
					comment = comment.Clear()
				case *Message:
					t.Comment = comment.AddSpace().Append(line.Comment).TrimSpace()
					p.Messages = append(p.Messages, t)
					comment = comment.Clear()
				case *Enum:
					t.Comment = comment.AddSpace().Append(line.Comment).TrimSpace()
					p.Enums = append(p.Enums, t)
					comment = comment.Clear()
				case *Package:
					t.Comment = comment.AddSpace().Append(line.Comment).TrimSpace()
					p.Name = t.Name
					p.Comment = comment.TrimSpace()
					comment = comment.Clear()
				case Comment:
					comment = comment.AddSpace().Append(t)
				default:
					Log.Debugf("Unhandled Return type for package: %T visitor\n", t)
				}
			}
		}
	}
	return nil
}
