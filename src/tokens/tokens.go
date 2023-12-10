package tokens

type TokenType string

type Token struct {
	Type TokenType
	Literal string
	Ln int
	Col int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	INCREMENT = "++"
	MINUS    = "-"
	DECREMENT = "--"
	BANG     = "!"
	ASTERISK = "*"
	EXPONENT = "**"
	SLASH    = "/"

	LT = "<"
	GT = ">"
	// todo: maybe support <= and >= operators
	// todo: maybe binary operators
	// todo: or, and operators

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var KeywordsMap = map[string]TokenType{
	"func": FUNCTION,
	"let": LET,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE, 
}