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
	"strings"
)

func NewEnumVisitor() *EnumVisitor {
	Log.Debug("Initializing EnumVisitor")
	out := &EnumVisitor{visitors: make([]Visitor, 0)}
	out.visitors = append(out.visitors,
		&CommentVisitor{},
		&EnumValueVisitor{})
	return out
}

type EnumVisitor struct {
	visitors []Visitor
}

func (ev *EnumVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "enum ") && in.Token == OpenBrace
}

func (ev *EnumVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	Log.Debugf("Visiting Enum: %d registered visitors\n", len(ev.visitors))
	fValues := in.SplitSyntax()
	out := NewEnum(namespace, fValues[1], in.Comment)

	var comment = Comment("")

	for scanner.Scan() {
		n := scanner.ReadLine()
		if strings.HasSuffix(n.Token, CloseBrace) {
			break
		}
		for _, visitor := range ev.visitors {
			if visitor.CanVisit(n) {
				rt := visitor.Visit(
					scanner,
					n,
					Join(Period, namespace, out.Name))

				switch t := rt.(type) {
				case *EnumValue:
					t.Comment = comment.AddSpace().Append(t.Comment).TrimSpace()
					out.Values = append(out.Values, t)
					comment = comment.Clear()
				case Comment:
					comment = comment.Append(t).AddSpace()
				default:
					Log.Infof("unable to parse enum value: %t", t)
				}
			}
		}
	}
	return out
}
