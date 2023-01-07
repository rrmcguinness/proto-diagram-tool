package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumValueVisitor_CanVisit(t *testing.T) {
	type args struct {
		in *Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test Enum Value", args: args{in: &Line{
			Syntax:  "RESIDENTIAL = 0",
			Token:   ";",
			Comment: "A residential address",
		}}, want: true},
		{name: "Test Not Enum Value", args: args{in: &Line{
			Syntax:  "message Address",
			Token:   "{",
			Comment: "Not an Enum",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evv := EnumValueVisitor{}
			assert.Equalf(t, tt.want, evv.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestEnumValueVisitor_Visit(t *testing.T) {
	type args struct {
		in0       Scanner
		in        *Line
		namespace string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Test Visitor", args: args{in0: nil, in: &Line{
			Syntax:  "RESIDENTIAL = 0",
			Token:   ";",
			Comment: "A residential address",
		}, namespace: "test"}, want: NewEnumValue("test", "0", "RESIDENTIAL", "A residential address")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evv := EnumValueVisitor{}
			assert.Equalf(t, tt.want, evv.Visit(tt.args.in0, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.in0, tt.args.in, tt.args.namespace)
		})
	}
}
