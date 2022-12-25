package proto

import "errors"

// Effective Final Variables

// InvalidImport = used during import sequence of files.
var InvalidImport = errors.New("invalid import")

// Log is the Package Logger
var Log = &Logger{}

var RegisteredVisitors []Visitor

// Initialize the visitors
func init() {
	// Handle Comments
	RegisteredVisitors = append(RegisteredVisitors, &CommentVisitor{})
	// Handle Package
	RegisteredVisitors = append(RegisteredVisitors, &PackageVisitor{})
	// Handle Imports
	RegisteredVisitors = append(RegisteredVisitors, &ImportVisitor{})
	// Handle Options
	RegisteredVisitors = append(RegisteredVisitors, &OptionVisitor{})
	// Handle Messages
	RegisteredVisitors = append(RegisteredVisitors, &MessageVisitor{})
	// Enums
	RegisteredVisitors = append(RegisteredVisitors, NewEnumVisitor())
	// Must be last as its patterns may encompass others.
	RegisteredVisitors = append(RegisteredVisitors, NewAttributeVisitor())
}
