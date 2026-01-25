// Package formula provides Excel function lexing and tokenization.
package formula

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

// LexerState represents the current state during lexing.
type LexerState int

const (
	StateText LexerState = iota
	StateNumber
	StateString
	StateIdentifier
	StateOperator
	StateParen
	StateComma
	StateColon
	StateFunction
	StateArgument
	StateArray
	StateQuoted
	StateError
)

// ExcelFunctionLexer handles Excel-specific function syntax tokenization.
type ExcelFunctionLexer struct {
	formula string
	pos     int
	tokens  []Token
	state   LexerState
	buffer  strings.Builder
}

// NewExcelFunctionLexer creates a new Excel function lexer.
func NewExcelFunctionLexer(formula string) *ExcelFunctionLexer {
	return &ExcelFunctionLexer{
		formula: formula,
		pos:     0,
		tokens:  make([]Token, 0),
		state:   StateText,
		buffer:  strings.Builder{},
	}
}

// TokenizeFunction lexes an Excel function string into tokens.
func (l *ExcelFunctionLexer) TokenizeFunction() ([]Token, error) {
	if !strings.HasPrefix(l.formula, "=") {
		return nil, fmt.Errorf("Excel function must start with '='")
	}

	// Remove leading '='
	funcBody := l.formula[1:]
	l.pos = 0
	l.state = StateFunction

	for l.pos < len(funcBody) {
		ch := rune(funcBody[l.pos])
		l.pos++

		switch l.state {
		case StateFunction:
			if unicode.IsLetter(ch) || ch == '_' {
				l.buffer.WriteRune(ch)
				l.state = StateIdentifier
			} else if unicode.IsDigit(ch) {
				l.buffer.WriteRune(ch)
				l.state = StateNumber
			} else if ch == '.' {
				if l.buffer.Len() > 0 && unicode.IsDigit(l.peekNext()) {
					l.buffer.WriteRune(ch)
					l.state = StateNumber // Decimal number
				} else {
					l.emitToken(TokenOperator, ".")
					l.buffer.Reset()
					l.state = StateText
				}
			} else if ch == ' ' {
				if l.buffer.Len() > 0 {
					// Function name complete, start of arguments
					l.emitFunctionToken()
					l.buffer.Reset()
					l.state = StateFunction
				} else {
					// Quoted argument
					l.state = StateQuoted
				}
			} else if unicode.IsSpace(ch) {
				if l.buffer.Len() > 0 {
					l.emitTokenBasedOnState()
					l.buffer.Reset()
				}
				l.state = StateText // Skip whitespace
			} else {
				l.buffer.WriteRune(ch)
				l.state = StateText
			}

		case StateQuoted:
			if ch == '"' {
				l.state = StateText // End quote
			} else {
				l.buffer.WriteRune(ch) // Quoted content
			}

		case StateIdentifier:
			if unicode.IsLetter(ch) || ch == '_' {
				l.buffer.WriteRune(ch)
			} else if unicode.IsDigit(ch) {
				l.buffer.WriteRune(ch)
				l.state = StateNumber
			} else if ch == ' ' {
				if l.buffer.Len() > 0 {
					// End function name, start of arguments
					l.emitToken(TokenFunction, l.buffer.String())
					l.buffer.Reset()
					l.state = StateFunction
				} else {
					// Quoted argument
					l.state = StateQuoted
				}
			} else if ch == '(' {
				l.emitToken(TokenFunction, l.buffer.String())
				l.emitToken(TokenParen, "(")
				l.buffer.Reset()
				l.state = StateText
			} else if unicode.IsSpace(ch) {
				if l.buffer.Len() > 0 {
					l.emitTokenBasedOnState()
					l.buffer.Reset()
				}
				l.state = StateText
			} else {
				l.buffer.WriteRune(ch)
			}

		case StateNumber:
			if unicode.IsDigit(ch) || ch == '.' {
				l.buffer.WriteRune(ch)
			} else if ch == 'e' || ch == 'E' {
				l.buffer.WriteRune(ch) // Scientific notation
			} else if unicode.IsSpace(ch) || ch == ',' || ch == ')' {
				l.emitToken(TokenNumber, l.buffer.String())
				l.buffer.Reset()
				l.state = StateText

				if ch == ',' {
					l.emitToken(TokenComma, ",")
				} else if ch == ')' {
					l.emitToken(TokenParen, ")")
				} else {
					l.state = StateText
				}
			} else {
				l.buffer.WriteRune(ch)
			}

		case StateArgument:
			if unicode.IsSpace(ch) || ch == ',' || ch == ')' {
				l.emitTokenBasedOnState()
				l.buffer.Reset()
				l.state = StateText

				if ch == ',' {
					l.emitToken(TokenComma, ",")
				} else if ch == ')' {
					l.emitToken(TokenParen, ")")
				} else {
					l.state = StateText
				}
			} else {
				l.buffer.WriteRune(ch)
			}

		case StateArray:
			if ch == '}' {
				l.state = StateText
			} else if unicode.IsSpace(ch) || ch == ',' || ch == ')' {
				l.emitTokenBasedOnState()
				l.buffer.Reset()
				l.state = StateText

				if ch == ',' {
					l.emitToken(TokenComma, ",")
				} else if ch == ')' {
					l.emitToken(TokenParen, ")")
				} else {
					l.state = StateText
				}
			} else {
				l.buffer.WriteRune(ch)
			}
		}
	}

	// Emit any remaining buffer content
	if l.buffer.Len() > 0 {
		l.emitTokenBasedOnState()
	}

	return l.tokens, nil
}

// emitTokenBasedOnState emits a token based on current lexer state.
func (l *ExcelFunctionLexer) emitTokenBasedOnState() {
	if l.buffer.Len() == 0 {
		return
	}

	var tokenType TokenType
	switch l.state {
	case StateIdentifier:
		tokenType = TokenFunction
	case StateNumber:
		tokenType = TokenLiteral
	case StateQuoted:
		tokenType = TokenLiteral
	case StateArray:
		tokenType = TokenLiteral
	default:
		tokenType = TokenLiteral
	}

	l.tokens = append(l.tokens, Token{
		Type:  tokenType,
		Value: l.buffer.String(),
	})
}

// emitFunctionToken emits a function token with name validation.
func (l *ExcelFunctionLexer) emitFunctionToken() {
	funcName := l.buffer.String()

	// Validate function name (must be valid Excel function name)
	if !l.isValidFunctionName(funcName) {
		l.tokens = append(l.tokens, Token{
			Type:  TokenError,
			Value: fmt.Sprintf("#NAME? Invalid function name: %s", funcName),
		})
		return
	}

	l.tokens = append(l.tokens, Token{
		Type:  TokenFunction,
		Value: funcName,
	})
}

// isValidFunctionName checks if a string is a valid Excel function name.
func (l *ExcelFunctionLexer) isValidFunctionName(name string) bool {
	if len(name) == 0 {
		return false
	}

	// Check against common Excel function names
	validFunctions := map[string]bool{
		// Math functions
		"SUM", "AVERAGE", "COUNT", "MAX", "MIN", "PRODUCT", "STDEV", "VAR",
		"ABS", "SQRT", "POWER", "LOG", "EXP", "ROUND", "CEILING", "FLOOR",
		// Text functions
		"CONCATENATE", "LEFT", "RIGHT", "MID", "LEN", "FIND", "SEARCH", "REPLACE",
		"SUBSTITUTE", "LOWER", "UPPER", "PROPER", "TRIM",
		// Logical functions
		"IF", "AND", "OR", "NOT", "TRUE", "FALSE", "ISERROR", "ISNUMBER", "ISTEXT",
		// Date functions
		"DATE", "TIME", "YEAR", "MONTH", "DAY", "TODAY", "NOW", "EDATE",
		// Lookup functions
		"VLOOKUP", "HLOOKUP", "INDEX", "MATCH", "CHOOSE",
		// Financial functions
		"PV", "FV", "PMT", "RATE", "NPER", "IRR", "NPV",
	}

	return validFunctions[strings.ToUpper(name)]
}

// peekNext returns the next character without advancing position.
func (l *ExcelFunctionLexer) peekNext() rune {
	if l.pos >= len(l.formula) {
		return 0
	}
	return rune(l.formula[l.pos])
}

// emitToken emits a token of the specified type.
func (l *ExcelFunctionLexer) emitToken(tokenType TokenType, value string) {
	l.tokens = append(l.tokens, Token{
		Type:  tokenType,
		Value: value,
	})
}

// GetTokens returns all lexed tokens.
func (l *ExcelFunctionLexer) GetTokens() []Token {
	return l.tokens
}
