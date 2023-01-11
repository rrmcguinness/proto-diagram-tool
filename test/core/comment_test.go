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

func TestComment_AddSpace(t *testing.T) {
	tests := []struct {
		name string
		c    proto.Comment
		want proto.Comment
	}{
		{name: "Add Space", c: proto.Comment("Test"), want: proto.Comment("Test ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.AddSpace(), "AddSpace()")
		})
	}
}

func TestComment_Append(t *testing.T) {
	type args struct {
		other proto.Comment
	}
	tests := []struct {
		name string
		c    proto.Comment
		args args
		want proto.Comment
	}{
		{name: "Append Comment", c: proto.Comment("Test"), args: args{other: proto.Comment(" test")}, want: proto.Comment("Test test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.Append(tt.args.other), "Append(%v)", tt.args.other)
		})
	}
}

func TestComment_Clear(t *testing.T) {
	tests := []struct {
		name string
		c    proto.Comment
		want proto.Comment
	}{
		{name: "Clear", c: proto.Comment("Test"), want: proto.Comment("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.Clear(), "Clear()")
		})
	}
}

func TestComment_ToMermaid(t *testing.T) {
	tests := []struct {
		name string
		c    proto.Comment
		want string
	}{
		{name: "To Mermaid", c: proto.Comment("Test"), want: "%% Test\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.ToMermaid(), "ToMermaid()")
		})
	}
}

func TestComment_TrimSpace(t *testing.T) {
	tests := []struct {
		name string
		c    proto.Comment
		want proto.Comment
	}{
		{name: "Trim Space", c: proto.Comment(" Test "), want: proto.Comment("Test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.TrimSpace(), "TrimSpace()")
		})
	}
}
