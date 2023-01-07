package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumVisitor_CanVisit(t *testing.T) {
	type fields struct {
		visitors []Visitor
	}
	type args struct {
		in *Line
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "Test Enum", fields: fields{
			visitors: []Visitor{},
		}, args: args{in: &Line{
			Syntax:  "enum AddressType",
			Token:   "{",
			Comment: "Enum Comment",
		}}, want: true},
		{name: "Test Bad Enum", fields: fields{
			visitors: []Visitor{},
		}, args: args{in: &Line{
			Syntax:  "message Address",
			Token:   "{",
			Comment: "Not an Enum",
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := &EnumVisitor{
				visitors: tt.fields.visitors,
			}
			assert.Equalf(t, tt.want, ev.CanVisit(tt.args.in), "CanVisit(%v)", tt.args.in)
		})
	}
}

func TestEnumVisitor_Visit(t *testing.T) {
	type fields struct {
		visitors []Visitor
	}

	type args struct {
		scanner   Scanner
		in        *Line
		namespace string
	}

	scanner := NewTestScanner("enum AddressType {")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "Test Visitor",
			fields: fields{visitors: []Visitor{}},
			args: args{
				scanner: scanner,
				in: &Line{
					Syntax:  "enum AddressType",
					Token:   "{",
					Comment: "Address Type",
				},
				namespace: "test",
			},
			want: NewEnum("test.AddressType", "AddressType", "Address Type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := &EnumVisitor{
				visitors: tt.fields.visitors,
			}
			assert.Equalf(t, tt.want, ev.Visit(tt.args.scanner, tt.args.in, tt.args.namespace), "Visit(%v, %v, %v)", tt.args.scanner, tt.args.in, tt.args.namespace)
		})
	}
}

func TestNewEnumVisitor(t *testing.T) {
	tests := []struct {
		name string
		want *EnumVisitor
	}{
		{name: "Test New Visitor", want: NewEnumVisitor()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEnumVisitor(), "NewEnumVisitor()")
		})
	}
}
