package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeFormat(t *testing.T) {
	type args struct {
		length int
		value  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Good Format", args: args{length: 10, value: "test"}, want: "test      |"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ComputeFormat(tt.args.length, tt.args.value), "ComputeFormat(%v, %v)", tt.args.length, tt.args.value)
		})
	}
}

func TestDashLine(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DashLine(tt.args.length), "DashLine(%v)", tt.args.length)
		})
	}
}

func TestMarkdownTable_AddHeader(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	type args struct {
		names []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add Header",
			args: args{names: []string{"test"}},
			fields: fields{
				header:        make([]string, 0),
				columnLengths: make([]int, 0),
				data:          make([][]string, 0),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				header:        tt.fields.header,
				columnLengths: tt.fields.columnLengths,
				data:          tt.fields.data,
			}
			mt.AddHeader(tt.args.names...)
			assert.Equal(t, 1, len(mt.header))
		})
	}
}

func TestMarkdownTable_EvaluateWidth(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	type args struct {
		i int
		d string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Evaluate Width",
			fields: fields{
				header:        make([]string, 0),
				columnLengths: make([]int, 0),
				data:          make([][]string, 0),
			}, args: args{
				i: 0,
				d: "test",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				header:        tt.fields.header,
				columnLengths: tt.fields.columnLengths,
				data:          tt.fields.data,
			}
			mt.EvaluateWidth(tt.args.i, tt.args.d)
			assert.Equal(t, 6, mt.columnLengths[0])
		})
	}
}

func TestMarkdownTable_Insert(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	type args struct {
		data []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Insert",
			fields: fields{
				header:        make([]string, 0),
				columnLengths: make([]int, 0),
				data:          make([][]string, 0),
			}, args: args{data: []string{"test1", "test2", "test3"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				header:        tt.fields.header,
				columnLengths: tt.fields.columnLengths,
				data:          tt.fields.data,
			}
			mt.Insert(tt.args.data...)
			assert.Equal(t, 3, len(mt.data[0]))
		})
	}
}

func TestMarkdownTable_String(t *testing.T) {
	type fields struct {
		header        []string
		columnLengths []int
		data          [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// The padding in this test is because it's not using
		// the Insert method.
		{
			name: "String Test",
			fields: fields{
				header:        []string{" c1"},
				columnLengths: []int{6},
				data:          [][]string{{" test"}},
			},
			want: "| c1   |\n|------|\n| test |\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &MarkdownTable{
				header:        tt.fields.header,
				columnLengths: tt.fields.columnLengths,
				data:          tt.fields.data,
			}
			assert.Equalf(t, tt.want, mt.String(), "String()")
		})
	}
}

func TestNewMarkdownTable(t *testing.T) {
	tests := []struct {
		name string
		want *MarkdownTable
	}{
		{name: "New Table", want: &MarkdownTable{
			header:        []string{},
			columnLengths: []int{},
			data:          [][]string{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewMarkdownTable(), "NewMarkdownTable()")
		})
	}
}
