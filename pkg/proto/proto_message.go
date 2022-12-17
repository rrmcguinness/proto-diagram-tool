package proto

import (
	"crypto"
	_ "crypto/md5"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	RegexAttributeMatcher = `^(\?|repeated \w+|map\<.*?>|\w+)\s+(.*?)\s+=\s+(\d+)(.*?);$`
	RegexMessageMatcher   = `message\s+(.*?)\s+{`
	RegexRepeated         = `(.*)repeated\s+(\w+)(.*)`
	RegexpMapMatcher      = `.*map\<(.*?),(.*?)\>.*`
)

var messageMatcher *regexp.Regexp
var attributeMatcher *regexp.Regexp
var repeatedMatcher *regexp.Regexp
var mapMatcher *regexp.Regexp

func init() {
	messageMatcher = regexp.MustCompile(RegexMessageMatcher)
	attributeMatcher = regexp.MustCompile(RegexAttributeMatcher)
	repeatedMatcher = regexp.MustCompile(RegexRepeated)
	mapMatcher = regexp.MustCompile(RegexpMapMatcher)
}

// Message represents a UML Class
type Message struct {
	ID         string
	FQN        string
	Name       string
	Comment    string
	attributes []*Attribute
}

func NewMessage(pkgName string, def []string, comment string) *Message {
	nLine := def[0]

	if messageMatcher.MatchString(nLine) {

		nGroups := messageMatcher.FindStringSubmatch(nLine)

		messageName := nGroups[1]
		messageFQN := pkgName + "." + messageName

		// Use a SHA 1 to generate the ID
		h := crypto.MD5.New()
		io.WriteString(h, messageFQN)

		out := Message{
			ID:         fmt.Sprintf("%x", h.Sum(nil)),
			FQN:        messageFQN,
			Name:       messageName,
			Comment:    comment,
			attributes: make([]*Attribute, 0),
		}

		for _, r := range def[1:] {
			r = strings.TrimSpace(r)
			if attributeMatcher.MatchString(r) {
				rVals := attributeMatcher.FindStringSubmatch(r)
				out.attributes = append(out.attributes, NewAttribute(rVals[2], rVals[1], rVals[3]))
			}
		}
		return &out
	}
	return nil
}

// Attribute A message attribute
type Attribute struct {
	Name  string
	Type  string
	Order int
}

func NewAttribute(name string, fqn string, order string) *Attribute {
	ord, _ := strconv.ParseInt(order, 10, 64)
	return &Attribute{
		Name:  name,
		Type:  fqn,
		Order: int(ord),
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("ID: %s; FQN: %s; Total Attributes: %d",
		m.ID, m.FQN, len(m.attributes))
}

func (m *Message) PlantUML() string {
	var out string
	out += fmt.Sprintf("struct .%s {\n", m.Name)
	for _, a := range m.attributes {
		out += fmt.Sprintf("  + %s %s\n", a.Type, a.Name)
	}
	out += "}"
	out = strings.ReplaceAll(out, repeated, brackets)
	return out
}

const stdTypes = "double,float,float32,float64,int,int32,int64,uint32,uint64,sint32,sint64,fixed32,fixed64,sfixed32,sfixed64,bool,string,bytes"
const repeated = "repeated "
const brackets = "[]"

func (m *Message) MermaidRelationships() string {
	dependencies := make([]string, 0)
	for _, a := range m.attributes {
		var compositeType []string

		t := a.Type
		if repeatedMatcher.MatchString(a.Type) {
			repGroup := repeatedMatcher.FindStringSubmatch(a.Type)
			compositeType = append(compositeType, repGroup[2])
			t = fmt.Sprintf("List~%s~", repGroup[2])
		}
		if mapMatcher.MatchString(a.Type) {
			mapGroup := mapMatcher.FindStringSubmatch(a.Type)
			compositeType = append(compositeType, mapGroup[1])
			compositeType = append(compositeType, mapGroup[2])
		}
		if !strings.Contains(stdTypes, t) {
			if len(compositeType) > 0 {
				for _, c := range compositeType {
					if !strings.Contains(stdTypes, c) {
						dependencies = append(dependencies, fmt.Sprintf(PatternRelationshipAggregation, m.Name, c))
					}
				}
			} else {
				dependencies = append(dependencies, fmt.Sprintf(PatternRelationshipDependency, m.Name, t))
			}
		}
	}
	return strings.Join(dependencies, "\n")
}

func (m *Message) Mermaid() string {
	var out string
	out += m.MermaidRelationships()
	out += fmt.Sprintf("\n  class %s{\n", m.Name)
	for _, a := range m.attributes {
		t := a.Type
		if repeatedMatcher.MatchString(a.Type) {
			repGroup := repeatedMatcher.FindStringSubmatch(a.Type)
			t = fmt.Sprintf("List~%s~", repGroup[2])
		}
		if mapMatcher.MatchString(a.Type) {
			mapGroup := mapMatcher.FindStringSubmatch(a.Type)
			t = fmt.Sprintf("Map~%s,%s~", mapGroup[1], mapGroup[2])
		}
		out += fmt.Sprintf("    +%s %s\n", t, a.Name)
	}
	out += "  }\n"
	return out
}
