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

import "fmt"

// Enum represents a Proto Enum type.
type Enum struct {
	*Qualified
	Values []*EnumValue
}

// NewEnum is the Enum Constructor
func NewEnum(q string, name string, comment Comment) *Enum {
	return &Enum{
		Qualified: &Qualified{
			Qualifier: q,
			Name:      name,
			Comment:   comment,
		},
		Values: make([]*EnumValue, 0),
	}
}

// ToMermaid prints a mermaid representation of the Enum
func (e Enum) ToMermaid() string {
	out := fmt.Sprintf("%s\nclass %s{\n  <<enumeration>>\n", e.Comment.ToMermaid(), e.Name)
	for _, v := range e.Values {
		out += fmt.Sprintf("  %s\n", v.Value)
	}
	out += "}"
	return out
}
