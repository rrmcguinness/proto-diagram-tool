package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEnumValue(t *testing.T) {
	type args struct {
		namespace string
		ordinal   string
		value     string
		comment   Comment
	}
	tests := []struct {
		name string
		args args
		want *EnumValue
	}{
		{name: "Test Enum Value", args: args{
			namespace: "test",
			ordinal:   "1",
			value:     "TEST",
			comment:   "Test",
		}, want: NewEnumValue("test", "1", "TEST", "Test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEnumValue(tt.args.namespace, tt.args.ordinal, tt.args.value, tt.args.comment), "NewEnumValue(%v, %v, %v, %v)", tt.args.namespace, tt.args.ordinal, tt.args.value, tt.args.comment)
		})
	}
}
