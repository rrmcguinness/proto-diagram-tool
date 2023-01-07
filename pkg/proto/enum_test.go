package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEnum(t *testing.T) {
	type args struct {
		q       string
		name    string
		comment Comment
	}
	tests := []struct {
		name string
		args args
		want *Enum
	}{
		{name: "Test Enum", args: args{q: "test", name: "TEST", comment: "Test"}, want: &Enum{
			Qualified: &Qualified{
				Qualifier: "test",
				Name:      "TEST",
				Comment:   "Test",
			},
			Values: []*EnumValue{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEnum(tt.args.q, tt.args.name, tt.args.comment), "NewEnum(%v, %v, %v)", tt.args.q, tt.args.name, tt.args.comment)
		})
	}
}
