package proto

// NamedValue is super class to capture names and values for typed lines.
type NamedValue struct {
	Name    string
	Value   any
	Comment Comment
}

// Qualified is a super class to capture namespace aware attributes and enums
type Qualified struct {
	Qualifier string
	Name      string
	Comment   Comment
}

/* Generalized Classes */

// SetDebug is used to enable the debug output, useful for troubleshooting.
func SetDebug(debug bool) {
	Log.debug = debug
}
