package lexer

import "github.com/varshard/monkey/token"

type Lexer struct {
	Position     int
	ReadPosition int
	Input        string
	Line         int
	Col          int
	EndCol       int
}

func New(input string) *Lexer {
	l := Lexer{
		Input: input,
		Line:  1,
	}

	return &l
}

func (l *Lexer) ReadChar() byte {
	char := l.Input[l.ReadPosition]
	if '\n' == char {
		l.Line += 1
		l.Col = 0
	}
	l.Position = l.ReadPosition
	l.ReadPosition += 1
	l.Col += 1
	return char
}

func (l *Lexer) NextToken() token.Token {
	if l.ReadPosition >= len(l.Input) {
		return token.Token{
			Position: l.ReadPosition,
			Type:     token.Eof,
			Line:     l.Line,
			Col:      l.Col,
		}
	}
	l.skipWhiteSpaces()
	char := l.ReadChar()
	tok := token.Token{
		Position: l.Position,
	}

	tokType, ok := tok.LookUpSymbol(char)
	if ok {
		tok.Type = tokType
		tok.Literal = string(char)
	} else if char == '=' {
		if l.peekChar() == '=' {
			l.ReadChar()
			tok.Literal = "=="
			tok.Type = token.Equal
		} else {
			tok.Literal = "="
			tok.Type = token.Assign
		}
	} else if char == '!' {
		if l.peekChar() == '=' {
			l.ReadChar()
			tok.Literal = "!="
			tok.Type = token.NotEqual
		} else {
			tok.Literal = "!"
			tok.Type = token.Bang
		}
	} else if char == '>' {
		if l.peekChar() == '=' {
			l.ReadChar()
			tok.Literal = ">="
			tok.Type = token.MoreThanEqual
		} else {
			tok.Literal = ">"
			tok.Type = token.MoreThan
		}
	} else if char == '<' {
		if l.peekChar() == '=' {
			l.ReadChar()
			tok.Literal = "<="
			tok.Type = token.LessThanEqual
		} else {
			tok.Literal = "<"
			tok.Type = token.LessThan
		}
	} else if char == '+' {
		if l.peekChar() == '+' {
			l.ReadChar()
			tok.Literal = "++"
			tok.Type = token.Increment
		} else {
			tok.Literal = "+"
			tok.Type = token.Plus
		}
	} else if char == '-' {
		if l.peekChar() == '-' {
			l.ReadChar()
			tok.Literal = "--"
			tok.Type = token.Decrement
		} else {
			tok.Literal = "-"
			tok.Type = token.Minus
		}
	} else if IsAlphabet(char) {
		tok = l.readIdentifier()
		// Handle let, if, else, and etc.
		tok.Type = tok.LookUpIdentifier(tok.Literal)
	} else if IsNumeric(char) {
		tok = l.readNumber()
	} else {
		tok.Literal = string(char)
		tok.Type = token.Illegal
	}
	// TODO: Read string

	tok.Line = l.Line
	tok.Col = l.Col
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.ReadPosition < len(l.Input) {
		return l.Input[l.ReadPosition]
	}

	return 0
}

func (l *Lexer) peekChars(n int) []byte {
	chars := make([]byte, 0)
	for i := 1; i < n && l.Position+i <= len(l.Input); i++ {
		chars = append(chars, l.Input[l.Position+i])
	}

	return chars
}

func IsAlphabet(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}

func IsAlphaNumeric(char byte) bool {
	return IsAlphabet(char) || IsNumeric(char)
}

func IsNumeric(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) readLet() token.Token {
	position := l.Position
	chars := l.peekChars(3)

	if string(chars) == "let" {
		return token.Token{
			Type:     token.Let,
			Literal:  "let",
			Position: position,
		}
	}

	return token.Token{}
}

func (l *Lexer) readIdentifier() token.Token {
	position := l.Position
	chars := []byte{l.Input[l.Position]}
	for {
		if !IsAlphaNumeric(l.peekChar()) {
			break
		}
		chars = append(chars, l.ReadChar())
	}

	return token.Token{
		Type:     token.Identifier,
		Literal:  string(chars),
		Position: position,
	}
}

func (l *Lexer) readNumber() token.Token {
	position := l.Position
	chars := []byte{l.Input[l.Position]}
	tokenType := token.Integer
	for {
		peekedChar := l.peekChar()

		if peekedChar == '.' {
			tokenType = token.Decimal
		}
		if !IsNumeric(peekedChar) && peekedChar != '.' {
			break
		}
		chars = append(chars, l.ReadChar())
	}

	return token.Token{
		Type:     tokenType,
		Literal:  string(chars),
		Position: position,
	}
}

func (l *Lexer) skipWhiteSpaces() {
	for {
		peekedChar := l.peekChar()
		if ' ' != peekedChar && '\t' != peekedChar && '\n' != peekedChar {
			break
		}
		l.ReadChar()
	}
}
