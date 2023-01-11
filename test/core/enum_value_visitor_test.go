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

func TestEnumValueVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *proto.Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test Enum Value", args: args{in: &proto.Line{
			Syntax:  "RESIDENTIAL = 0",
			Token:   ";",
			Comment: "A residential address",
		}}, want: true},
		{name: "Test Not Enum Value", args: args{in: &proto.Line{
			Syntax:  "message Address",
			Token:   "{",
			Comment: "Not an Enum",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evv := proto.EnumValueVisitor{}
			assert.Equalf(t, tt.want, evv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestEnumValueVisitor_Visit(t *testing.T) {
	type args struct {
		in0       proto.Scanner
		in        *proto.Line
		namespace string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Test Visitor", args: args{in0: nil, in: &proto.Line{
			Syntax:  "RESIDENTIAL = 0",
			Token:   ";",
			Comment: "A residential address",
		}, namespace: "test"}, want: proto.NewEnumValue("test", "0", "RESIDENTIAL", "A residential address")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evv := proto.EnumValueVisitor{}
			assert.Equalf(t, tt.want, evv.Visit(tt.args.in0, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.namespace)
		})
	}
}
