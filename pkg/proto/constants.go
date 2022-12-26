package proto

const (
	Repeated = "repeated"
	Map      = "map"
	Reserved = "reserved"

	SpaceRemovalRegex = `\s+`
	Period            = "."
	Empty             = ""
	Space             = " "
	OpenBrace         = "{"
	CloseBrace        = "}"
	OpenBracket       = "["
	ClosedBracket     = "]"
	Semicolon         = ";"
	Comma             = ","

	InlineCommentPrefix        = "//"
	MultiLineCommentInitiator  = "/*"
	MultilineCommentTerminator = "*/"
	OpenMap                    = "map<"
	CloseMap                   = ">"
	DoubleQuote                = `"`
	SingleQuote                = `'`
)
