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

func TestMessage_HasAttributes(t *testing.T) {
	type fields struct {
		Attributes []*proto.Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Has Attribute",
			fields: fields{Attributes: []*proto.Attribute{
				proto.NewAttribute("test", "None"),
			}},
			want: true},
		{name: "Doesn't Have Attributes",
			fields: fields{Attributes: make([]*proto.Attribute, 0)},
			want:   false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &proto.Message{
				Attributes: tt.fields.Attributes,
			}
			assert.Equalf(t, tt.want, m.HasAttributes(), "HasAttributes()")
		})
	}
}

func TestMessage_HasEnums(t *testing.T) {
	type fields struct {
		Enums []*proto.Enum
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Has Enums",
			fields: fields{
				Enums: []*proto.Enum{proto.NewEnum("test", "test", "test")}},
			want: true},
		{name: "Doesn't Have Enums",
			fields: fields{
				Enums: make([]*proto.Enum, 0)},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &proto.Message{
				Enums: tt.fields.Enums,
			}
			assert.Equalf(t, tt.want, m.HasEnums(), "HasEnums()")
		})
	}
}

func TestMessage_HasMessages(t *testing.T) {
	type fields struct {
		Messages []*proto.Message
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Has Messages", fields: fields{Messages: []*proto.Message{proto.NewMessage()}},
			want: true},
		{name: "Doesn't have Messages", fields: fields{Messages: make([]*proto.Message, 0)},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &proto.Message{
				Messages: tt.fields.Messages,
			}
			assert.Equalf(t, tt.want, m.HasMessages(), "HasMessages()")
		})
	}
}

func TestNewMessage(t *testing.T) {
	tests := []struct {
		name string
		want *proto.Message
	}{
		{name: "New Message", want: &proto.Message{
			Qualified:  &proto.Qualified{},
			Attributes: make([]*proto.Attribute, 0),
			Messages:   make([]*proto.Message, 0),
			Enums:      make([]*proto.Enum, 0),
			Reserved:   make([]*proto.Reserved, 0),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, proto.NewMessage(), "NewMessage()")
		})
	}
}
