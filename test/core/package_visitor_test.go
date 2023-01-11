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

func TestPackageVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *proto.Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Can Visit", args: args{in: &proto.Line{
			Syntax:  "package test.test",
			Token:   ";",
			Comment: "// Test Comment",
		}}, want: true},
		{name: "Can't Visit", args: args{in: &proto.Line{Token: "//"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pv := &proto.PackageVisitor{}
			assert.Equalf(t, tt.want, pv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestPackageVisitor_Visit(t *testing.T) {
	type args struct {
		in0 proto.Scanner
		in  *proto.Line
		in2 string
	}
	testReader := NewTestScanner(``)

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Positive Visit", args: args{
			in0: testReader,
			in: &proto.Line{
				Syntax:  "package test.package",
				Token:   ";",
				Comment: "Test Package",
			},
			in2: "test",
		}, want: &proto.Package{
			Path:     "",
			Name:     "test.package",
			Comment:  "Test Package",
			Options:  make([]*proto.Option, 0),
			Imports:  make([]*proto.Import, 0),
			Messages: make([]*proto.Message, 0),
			Enums:    make([]*proto.Enum, 0),
			Services: make([]*proto.Service, 0),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pv := &proto.PackageVisitor{}
			assert.Equalf(t, tt.want, pv.Visit(tt.args.in0, tt.args.in, tt.args.in2), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.in2)
		})
	}
}
