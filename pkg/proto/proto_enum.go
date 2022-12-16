package proto

import (
	"crypto"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	RegexEnumMatcher          = `enum\s+(\w+)\s+{`
	RegexEnumAttributeMatcher = `\s+(.*?)\s+=\s+(\d+)([;|;\s+\/\/.*])`
)

var enumMatcher *regexp.Regexp
var enumAttributeMatcher *regexp.Regexp

func init() {
	enumMatcher = regexp.MustCompile(RegexEnumMatcher)
	enumAttributeMatcher = regexp.MustCompile(RegexEnumAttributeMatcher)
}

type EnumValue struct {
	Ordinal int
	Name    string
	Comment string
}

type Enum struct {
	ID      string
	FQN     string
	Name    string
	Comment string
	values  []*EnumValue
}

func NewEnum(pkg string, def []string, comment string) *Enum {
	if len(def) > 0 {
		eName := def[0]
		var out *Enum
		if enumMatcher.MatchString(eName) {
			enumName := enumMatcher.FindStringSubmatch(eName)[1]
			enumFQN := pkg + "." + enumName

			h := crypto.MD5.New()
			io.WriteString(h, enumFQN)

			out = &Enum{
				ID:      fmt.Sprintf("%x", h.Sum(nil)),
				FQN:     enumFQN,
				Name:    enumName,
				Comment: comment,
				values:  make([]*EnumValue, 0),
			}

			if len(def) > 1 {
				for _, v := range def[1:] {
					if enumAttributeMatcher.MatchString(v) {
						aGroups := enumAttributeMatcher.FindStringSubmatch(v)
						ord, _ := strconv.ParseInt(aGroups[2], 10, 64)
						cmt := aGroups[3]
						if strings.Contains(cmt, "//") {
							cmt = cmt[strings.Index(cmt, "//")+2:]
						} else {
							cmt = ""
						}
						out.values = append(out.values, &EnumValue{
							Ordinal: int(ord),
							Name:    aGroups[1],
							Comment: cmt,
						})
					}
				}
			}
		}
		return out
	}
	return nil
}

func (e *Enum) String() string {
	return fmt.Sprintf("ID: %s; FQN: %s; Total Values: %d",
		e.ID, e.FQN, len(e.values))
}

func (e *Enum) PlantUML() string {
	var out string
	out += fmt.Sprintf("enum .%s{\n", e.Name)
	for _, v := range e.values {
		out += fmt.Sprintf("  %s\n", v.Name)
	}
	out += "}\n"
	return out
}
