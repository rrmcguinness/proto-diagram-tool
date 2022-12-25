package proto

import (
	"strings"
)

var isDebug bool

type PackageVisitor struct {
}

func (pv *PackageVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "package ") && in.Token == Semicolon
}

func (pv *PackageVisitor) Visit(_ Scanner, in *Line, _ string) interface{} {
	fValues := in.SplitSyntax()
	return &Package{Name: RemoveSemicolon(fValues[1])}
}
