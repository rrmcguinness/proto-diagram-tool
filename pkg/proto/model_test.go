package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamedValue_GetAnchor(t *testing.T) {
	type fields struct {
		Name    string
		Value   string
		Comment Comment
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get Anchor", fields{
			Name:    "SomeName",
			Value:   "Some Value",
			Comment: "Some Comment",
		}, "some_name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			namedValue := &NamedValue{
				Name:    tt.fields.Name,
				Value:   tt.fields.Value,
				Comment: tt.fields.Comment,
			}
			assert.Equalf(t, tt.want, namedValue.GetAnchor(), "GetAnchor()")
		})
	}
}
