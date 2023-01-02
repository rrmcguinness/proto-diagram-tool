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

package proto

import (
	"fmt"
	"strings"
)

type Service struct {
	*Qualified
	Methods []*Rpc
}

func NewService(namespace string, name string, comment Comment) *Service {
	return &Service{
		Qualified: &Qualified{
			Qualifier: namespace,
			Name:      name,
			Comment:   comment,
		},
		Methods: make([]*Rpc, 0),
	}
}

func (s *Service) AddRpc(rpc ...*Rpc) {
	s.Methods = append(s.Methods, rpc...)
}

func CleanParameterName(in string) string {
	if strings.Contains(in, Period) {
		return in[strings.LastIndex(in, Period)+1:]
	} else {
		return in
	}
}

func FormatParameters(in []*Parameter) string {
	out := ""
	for i := 0; i < len(in); i++ {
		p := in[i]
		if p.Stream {
			out += fmt.Sprintf("Stream~%s~", CleanParameterName(p.Type))
		} else {
			out += CleanParameterName(in[i].Type)
		}
		if i < len(in)-1 {
			out += ","
		}
	}
	return out
}

func FormatRelationships(name string, in []*Parameter) string {
	out := ""
	for _, i := range in {
		t := strings.TrimSpace(i.Type)
		if strings.HasSuffix(t, name) {
			t = CleanParameterName(t)
		}

		if i.Stream {
			out += fmt.Sprintf("%s --o `%s`\n", name, t)
		} else {
			out += fmt.Sprintf("%s --> `%s`\n", name, t)
		}
	}
	return out
}

func (s *Service) ToMermaid() string {
	relationships := ""

	out := fmt.Sprintf("class %s {\n  <<service>>\n", s.Name)
	for _, m := range s.Methods {
		out += fmt.Sprintf("  +%s(%s) %s\n",
			m.Name,
			FormatParameters(m.InputParameters),
			FormatParameters(m.ReturnParameters))

		relationships += FormatRelationships(s.Name, m.InputParameters)
		relationships += FormatRelationships(s.Name, m.ReturnParameters)
	}
	out += "}\n"
	out += relationships

	return out
}
