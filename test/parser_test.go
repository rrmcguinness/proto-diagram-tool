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

import "testing"

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

}
