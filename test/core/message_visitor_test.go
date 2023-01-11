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

package core

import (
	"testing"

	"github.com/rrmcguinness/proto-diagram-tool/pkg/proto"
	"github.com/stretchr/testify/assert"
)

func TestMessageVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *proto.Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Can Visit", args: args{in: &proto.Line{
			Syntax:  "message Test",
			Token:   "{",
			Comment: "Test Message",
		}}, want: true},
		{name: "Can not Visit", args: args{in: &proto.Line{
			Token:   "//",
			Comment: "Test Message",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv := &proto.MessageVisitor{}
			assert.Equalf(t, tt.want, mv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestMessageVisitor_Visit(t *testing.T) {
	type args struct {
		scanner   proto.Scanner
		in        *proto.Line
		namespace string
	}
	testFile := `
  enum TestEnum {
    T1 = 0;
    T2 = 1;
  }
  string name = 1; // Name
  TestEnum type = 2; // Type
`
	scanner := NewTestScanner(testFile)

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Message Scanner", args: args{
			scanner:   scanner,
			in:        proto.NewLine("message Test { // Test Message"),
			namespace: "test",
		}, want: &proto.Message{
			Qualified: &proto.Qualified{
				Qualifier: "test.Test",
				Name:      "Test",
				Comment:   "Test Message",
			},
			Attributes: []*proto.Attribute{
				{
					Qualified: &proto.Qualified{
						Qualifier: "test.Test",
						Name:      "name",
						Comment:   "Name",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"string"},
					Ordinal:     1,
					Annotations: make([]*proto.Annotation, 0),
				},
				{
					Qualified: &proto.Qualified{
						Qualifier: "test.Test",
						Name:      "type",
						Comment:   "Type",
					},
					Repeated:    false,
					Map:         false,
					Kind:        []string{"TestEnum"},
					Ordinal:     2,
					Annotations: make([]*proto.Annotation, 0),
				},
			},
			Messages: make([]*proto.Message, 0),
			Enums: []*proto.Enum{
				{
					Qualified: &proto.Qualified{
						Qualifier: "test.Test.TestEnum",
						Name:      "TestEnum",
					},
					Values: []*proto.EnumValue{
						{
							Namespace: "test.Test.TestEnum",
							Ordinal:   0,
							Value:     "T1",
						},
						{
							Namespace: "test.Test.TestEnum",
							Ordinal:   1,
							Value:     "T2",
						},
					},
				},
			},
			Reserved: make([]*proto.Reserved, 0),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv := &proto.MessageVisitor{}
			assert.Equalf(t, tt.want, mv.Visit(tt.args.scanner, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.scanner, tt.args.in, tt.args.namespace)
		})
	}
}
