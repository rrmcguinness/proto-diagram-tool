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

package proto

import (
	"strings"
)

const (
	MarkdownPadding = 2
)

type MarkdownTable struct {
	header        []string
	columnLengths []int
	data          [][]string
}

func NewMarkdownTable() *MarkdownTable {
	return &MarkdownTable{header: make([]string, 0), columnLengths: make([]int, 0), data: make([][]string, 0)}
}

func (mt *MarkdownTable) EvaluateWidth(i int, d string) {
	dLen := len(d) + MarkdownPadding
	if len(mt.columnLengths) == i {
		mt.columnLengths = append(mt.columnLengths, dLen)
	} else if mt.columnLengths[i] < dLen {
		mt.columnLengths[i] = dLen
	}
}

func (mt *MarkdownTable) AddHeader(names ...string) {
	for i, d := range names {
		mt.EvaluateWidth(i, d)
		names[i] = Space + d
	}
	mt.header = append(mt.header, names...)
}

func (mt *MarkdownTable) Insert(data ...string) {
	for i, d := range data {
		mt.EvaluateWidth(i, d)
		// Pad
		data[i] = Space + d
	}
	mt.data = append(mt.data, data)
}

func ComputeFormat(length int, value string) string {
	out := value
	for i := 0; i < length-len(value); i++ {
		out += Space
	}
	out += Pipe
	return out
}

func DashLine(length int) string {
	return strings.Repeat(Hyphen, length) + Pipe
}

func (mt *MarkdownTable) String() string {
	// Write the Header
	out := Pipe
	for i, h := range mt.header {
		out += ComputeFormat(mt.columnLengths[i], h)
	}
	// Write the Header Separator
	out += EndL + Pipe
	for i, _ := range mt.header {
		out += DashLine(mt.columnLengths[i])
	}
	out += EndL
	// Write the data
	for i := 0; i < len(mt.data); i++ {
		out += Pipe
		for j := 0; j < len(mt.data[i]); j++ {
			out += ComputeFormat(mt.columnLengths[j], mt.data[i][j])
		}
		out += EndL
	}
	return out
}
