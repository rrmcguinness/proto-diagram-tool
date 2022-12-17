package proto

import (
	"fmt"
	"regexp"
)

const (
	RegexPackageMatcher = `^package\s+(.*?);$`
)

var patternPackage *regexp.Regexp

func init() {
	patternPackage = regexp.MustCompile(RegexPackageMatcher)
}

// Package the proto package, not the language package
type Package struct {
	Name     string
	Comments string
	Messages map[string]*Message
	Enums    map[string]*Enum
}

func IsPackage(in string) bool {
	return patternPackage.MatchString(in)
}

func GetPackageName(in string) string {
	groups := patternPackage.FindStringSubmatch(in)
	return groups[1]
}

func NewPackage(in string, comments string) *Package {
	return &Package{
		Name:     GetPackageName(in),
		Comments: comments,
		Messages: make(map[string]*Message),
		Enums:    make(map[string]*Enum),
	}
}

func (p *Package) AddMessage(message *Message) {
	p.Messages[message.FQN] = message
}

func (p *Package) AddEnum(e *Enum) {
	p.Enums[e.FQN] = e
}

func (p *Package) PlantUML() string {
	var out string
	out += fmt.Sprintf("package %s {\n", p.Name)
	for _, e := range p.Enums {
		out += fmt.Sprintf("%s\n", e.PlantUML())
	}
	for _, m := range p.Messages {
		out += fmt.Sprintf("%s\n", m.PlantUML())
	}
	out += "}\n"
	return out
}

func (p *Package) Mermaid() string {
	var out string
	out += "classDiagram\n"
	for _, m := range p.Messages {
		out += m.Mermaid()
	}

	for _, e := range p.Enums {
		out += e.Mermaid()
	}

	return out
}
