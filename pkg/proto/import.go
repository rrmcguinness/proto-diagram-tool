package proto

// Import represents an importable file
type Import struct {
	Path    string
	Comment Comment
}

// NewImport is the import constructor
func NewImport(path string) *Import {
	return &Import{Path: path}
}
