package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test import visitor CanVisit",
			args: args{in: NewLine("import \"google/protobuf/timestamp.proto\";")},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iv := &ImportVisitor{}
			assert.Equalf(t, tt.want, iv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestImportVisitor_Visit(t *testing.T) {
	type args struct {
		in0 Scanner
		in  *Line
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
				in: &Line{
					Syntax:  "import \"google/protobuf/timestamp.proto\"",
					Token:   ";",
					Comment: "",
				},
				in2: "",
			},
			want: &Import{
				Path:    "google/protobuf/timestamp.proto",
				Comment: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iv := &ImportVisitor{}
			assert.Equalf(t, tt.want, iv.Visit(tt.args.in0, tt.args.in, tt.args.in2), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.in2)
		})
	}
}
