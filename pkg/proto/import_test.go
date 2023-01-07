package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewImport(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Import
	}{
		{name: "Create", args: args{path: "com.google.test"}, want: &Import{
			Path:    "com.google.test",
			Comment: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewImport(tt.args.path), "NewImport(%v)", tt.args.path)
		})
	}
}
