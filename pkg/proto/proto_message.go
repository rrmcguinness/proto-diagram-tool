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
	RegexAttributeMatcher = `(\w+|map\<.*?\>|repeated\s+\w+)\s+(.*?)\s+=\s+(\d+)(.*?);`
	RegexMessageMatcher   = `message\s+(.*?)\s+{`
	RegexRepeated         = `repeated\s+(\w+)(.*)`
)

var messageMatcher *regexp.Regexp
var attributeMatcher *regexp.Regexp
var repeatedMatcher *regexp.Regexp

func init() {
	messageMatcher = regexp.MustCompile(RegexMessageMatcher)
	attributeMatcher = regexp.MustCompile(RegexAttributeMatcher)
	repeatedMatcher = regexp.MustCompile(RegexRepeated)
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
	out = strings.ReplaceAll(out, "repeated ", "[]")
	return out
}
