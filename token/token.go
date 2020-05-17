package token

type TokenType string

const (
	Illegal TokenType = "illegal"
	Eof     TokenType = "Eof"

	Lparen    TokenType = "("
	Rparen    TokenType = ")"
	Lbrace    TokenType = "{"
	Rbrace    TokenType = "}"
	Semicolon TokenType = ";"
	Comma     TokenType = ","
	Period    TokenType = "."

	Equal    TokenType = "=="
	NotEqual TokenType = "!="

	// Operators
	Plus     TokenType = "+"
	Minus    TokenType = "-"
	Multiply TokenType = "*"
	Divide   TokenType = "/"
	Assign   TokenType = "="
	Not      TokenType = "!"

	Let      TokenType = "let"
	Function TokenType = "fn"

	If   TokenType = "if"
	Else TokenType = "else"

	Identifier TokenType = "identifier"
	Integer    TokenType = "integer"
)

var reserved = map[string]TokenType{
	"fn":   Function,
	"let":  Let,
	"if":   If,
	"else": Else,
}

var symbols = map[byte]TokenType{
	'+': Plus,
	'-': Minus,
	'*': Multiply,
	'/': Divide,
	'{': Lbrace,
	'}': Rbrace,
	'(': Lparen,
	')': Rparen,
	',': Comma,
	'.': Period,
	';': Semicolon,
}

type Token struct {
	Col      int
	Line     int
	Type     TokenType
	Literal  string
	Position int
}

func (t Token) LookUpIdentifier(literal string) TokenType {
	tokType, ok := reserved[literal]
	if ok {
		return tokType
	}

	return Identifier
}

func (t Token) LookUpSymbol(char byte) (TokenType, bool) {
	tokType, ok := symbols[char]
	return tokType, ok
}