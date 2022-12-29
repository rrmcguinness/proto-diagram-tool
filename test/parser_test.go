/*
 * Copyright 2022 Google, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

type Parser struct {
	ignore     []string
	keywords   []string
	knownTypes []string
	tokens     []string
}

//Proto3Types = "double,float,int32,int64,uint32,uint64,sint32,sint64,fixed32,fixed64,sfixed32,sfixed64,bool,string,bytes"

// InitKnownTypes initializes known types
func (p *Parser) InitKnownTypes() {
	p.knownTypes = make([]string, 0)
	p.knownTypes = append(p.knownTypes,
		"double", "float", "int32",
		"uint32", "sint32", "fixed32",
		"int64", "uint64", "sint64",
		"fixed64", "bool", "string", "bytes")
}

func (p *Parser) InitKeyWords() {
	p.keywords = make([]string, 0)
	p.keywords = append(p.keywords, "syntax", "package", "import", "reserved", "to", "enum", "to max", "repeated")
}

func (p *Parser) InitTokens() {
	p.tokens = make([]string, 0)
	p.tokens = append(p.tokens, "}", "=", ";", ",")
}

func (p Parser) init() {
	p.InitKnownTypes()

	p.ignore = make([]string, 0)

}

func TestParser(t *testing.T) {

	cleaner := regexp.MustCompile(`\s+|\n`)

	dir := "protos/test/location"

	lines := make([]string, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Visiting Path: %s\n", path)

		if strings.HasSuffix(path, ".proto") {
			readFile, err := os.Open(path)
			if err != nil {
				return err
			}

			scanner := bufio.NewScanner(readFile)
			scanner.Split(bufio.ScanRunes)

			line := ""
			tokenReached := false

			for scanner.Scan() {
				rune := scanner.Text()
				if rune == ";" || rune == "}" || rune == "{" {
					lines = append(lines, cleaner.ReplaceAllString(strings.TrimSpace(line+rune), " "))
					tokenReached = true
					line = ""
				} else if strings.HasPrefix(line, "//") && rune == "\n" {
					if tokenReached {
						lines = append(lines, cleaner.ReplaceAllString(strings.TrimSpace(line), " "))
						pLine := lines[len(lines)-2]
						cLine := lines[len(lines)-1]
						lines[len(lines)-2] = cLine
						lines[len(lines)-1] = pLine
					} else {
						lines = append(lines, cleaner.ReplaceAllString(strings.TrimSpace(line), " "))
					}
					line = ""
					tokenReached = false
				} else if strings.HasPrefix(line, "/*") && strings.HasSuffix(line, "*/") {
					lines = append(lines, cleaner.ReplaceAllString(line, " "))
					line = ""
				} else {
					if rune != "\n" {
						if rune == " " {
							if len(line) > 0 {
								line += rune
							}
						} else {
							line += rune
						}
					} else {
						tokenReached = false
					}
				}
			}
		}
		return nil
	})

	for _, l := range lines {
		fmt.Println(l)
	}

	if err != nil {
		fmt.Printf("Error reading files: %v", err)
	}

}
