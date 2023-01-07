package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine_SplitSyntax(t *testing.T) {
	type fields struct {
		Syntax  string
		Token   string
		Comment Comment
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{name: "Test Split Syntax", fields: fields{
			Syntax:  "message AddressType",
			Token:   ";",
			Comment: "Test",
		}, want: []string{"message", "AddressType"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				Syntax:  tt.fields.Syntax,
				Token:   tt.fields.Token,
				Comment: tt.fields.Comment,
			}
			assert.Equalf(t, tt.want, l.SplitSyntax(), "SplitSyntax()")
		})
	}
}

func TestNewLine(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want *Line
	}{
		{name: "Test Comment", args: args{in: "// Comment"}, want: &Line{Token: "//", Comment: "Comment"}},
		{name: "Test Multiline Comment", args: args{in: "/* Comment */"}, want: &Line{Token: "/*", Comment: "Comment"}},
		{name: "Test Open Brace", args: args{in: "message AddressType { // Comment"}, want: &Line{Token: "{", Syntax: "message AddressType", Comment: "Comment"}},
		{name: "Test Semicolon", args: args{in: "string name = 1; // Comment"}, want: &Line{Token: ";", Syntax: "string name = 1", Comment: "Comment"}},
		{name: "Test Close Brace", args: args{in: "} // Comment"}, want: &Line{Token: "}", Syntax: "", Comment: "Comment"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewLine(tt.args.in), "NewLine(%v)", tt.args.in)
		})
	}
}
