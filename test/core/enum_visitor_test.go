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

func TestEnumVisitor_CanVisit(t *testing.T) {
	type fields struct {
		visitors []proto.Visitor
	}
	type args struct {
		in *proto.Line
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "Test Enum", fields: fields{
			visitors: []proto.Visitor{},
		}, args: args{in: &proto.Line{
			Syntax:  "enum AddressType",
			Token:   "{",
			Comment: "Enum Comment",
		}}, want: true},
		{name: "Test Bad Enum", fields: fields{
			visitors: []proto.Visitor{},
		}, args: args{in: &proto.Line{
			Syntax:  "message Address",
			Token:   "{",
			Comment: "Not an Enum",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := &proto.EnumVisitor{
				Visitors: tt.fields.visitors,
			}
			assert.Equalf(t, tt.want, ev.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestEnumVisitor_Visit(t *testing.T) {
	type fields struct {
		visitors []proto.Visitor
	}

	type args struct {
		scanner   proto.Scanner
		in        *proto.Line
		namespace string
	}

	scanner := NewTestScanner("enum AddressType {")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "Test Visitor",
			fields: fields{visitors: []proto.Visitor{}},
			args: args{
				scanner: scanner,
				in: &proto.Line{
					Syntax:  "enum AddressType",
					Token:   "{",
					Comment: "Address Type",
				},
				namespace: "test",
			},
			want: proto.NewEnum("test.AddressType", "AddressType", "Address Type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := &proto.EnumVisitor{
				Visitors: tt.fields.visitors,
			}
			assert.Equalf(t, tt.want, ev.Visit(tt.args.scanner, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.scanner, tt.args.in, tt.args.namespace)
		})
	}
}

func TestNewEnumVisitor(t *testing.T) {
	tests := []struct {
		name string
		want *proto.EnumVisitor
	}{
		{name: "Test New Visitor", want: proto.NewEnumVisitor()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, proto.NewEnumVisitor(), "NewEnumVisitor()")
		})
	}
}
