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

func TestImportVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *proto.Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test import visitor CanVisit",
			args: args{in: proto.NewLine("import \"google/protobuf/timestamp.proto\";")},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iv := &proto.ImportVisitor{}
			assert.Equalf(t, tt.want, iv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestImportVisitor_Visit(t *testing.T) {
	type args struct {
		in0 proto.Scanner
		in  *proto.Line
		in2 string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Test import Visitor",
			args: args{
				in0: NewTestScanner("import \"google/protobuf/timestamp.proto\";"),
				in: &proto.Line{
					Syntax:  "import \"google/protobuf/timestamp.proto\"",
					Token:   ";",
					Comment: "",
				},
				in2: "",
			},
			want: &proto.Import{
				Path:    "google/protobuf/timestamp.proto",
				Comment: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iv := &proto.ImportVisitor{}
			assert.Equalf(t, tt.want, iv.Visit(tt.args.in0, tt.args.in, tt.args.in2), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.in2)
		})
	}
}
