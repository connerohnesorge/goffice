//nolint:all // formula package is scaffolding for future formula evaluation - WIP
// Package formula provides formula parsing and evaluation for spreadsheets.
package formula

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

// Parser parses Excel formulas into AST expressions.
type Parser struct {
	tokens []spreadsheet.Token
	pos    int
}

// NewParser creates a new formula parser.
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a formula string into an expression.
func (p *Parser) Parse(formula string) (Expression, error) {
	// Tokenize the formula
	tokens, err := spreadsheet.TokenizeFormula(formula)
	if err != nil {
		return nil, err
	}

	p.tokens = tokens
	p.pos = 0

	if len(p.tokens) == 0 {
		return nil, fmt.Errorf("empty formula")
	}

	return p.parseComparison()
}

// current returns the current token, or nil if at end.
func (p *Parser) current() *spreadsheet.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}

	return &p.tokens[p.pos]
}

// peek returns the next token without advancing, or nil if at end.
func (p *Parser) peek() *spreadsheet.Token {
	if p.pos+1 >= len(p.tokens) {
		return nil
	}

	return &p.tokens[p.pos+1]
}

// advance moves to the next token.
func (p *Parser) advance() {
	if p.pos < len(p.tokens) {
		p.pos++
	}
}

// expect checks if the current token matches the expected type and value.
func (p *Parser) expect(tokenType spreadsheet.TokenType, value string) error {
	tok := p.current()
	if tok == nil {
		return fmt.Errorf("unexpected end of formula, expected %s", value)
	}
	if tok.Type != tokenType || tok.Value != value {
		return fmt.Errorf("expected %s, got %s", value, tok.Value)
	}

	return nil
}

// parseComparison parses comparison operators (=, <>, <, <=, >, >=).
// This is the lowest precedence level.
func (p *Parser) parseComparison() (Expression, error) {
	left, err := p.parseConcatenation()
	if err != nil {
		return nil, err
	}

	for {
		tok := p.current()
		if tok == nil {
			break
		}

		if tok.Type == spreadsheet.TokenOperator {
			switch tok.Value {
			case "=", "<>", "<", "<=", ">", ">=":
				op := tok.Value
				p.advance()
				right, err := p.parseConcatenation()
				if err != nil {
					return nil, err
				}
				left = &BinaryOp{Left: left, Right: right, Operator: op}
			default:
				return left, nil
			}
		} else {
			break
		}
	}

	return left, nil
}

// parseConcatenation parses string concatenation (&).
func (p *Parser) parseConcatenation() (Expression, error) {
	left, err := p.parseAddition()
	if err != nil {
		return nil, err
	}

	for {
		tok := p.current()
		if tok == nil {
			break
		}

		if tok.Type == spreadsheet.TokenOperator && tok.Value == "&" {
			p.advance()
			right, err := p.parseAddition()
			if err != nil {
				return nil, err
			}
			left = &BinaryOp{Left: left, Right: right, Operator: "&"}
		} else {
			break
		}
	}

	return left, nil
}

// parseAddition parses addition and subtraction (+, -).
func (p *Parser) parseAddition() (Expression, error) {
	left, err := p.parseMultiplication()
	if err != nil {
		return nil, err
	}

	for {
		tok := p.current()
		if tok == nil {
			break
		}

		if tok.Type == spreadsheet.TokenOperator && (tok.Value == "+" || tok.Value == "-") {
			op := tok.Value
			p.advance()
			right, err := p.parseMultiplication()
			if err != nil {
				return nil, err
			}
			left = &BinaryOp{Left: left, Right: right, Operator: op}
		} else {
			break
		}
	}

	return left, nil
}

// parseMultiplication parses multiplication and division (*, /).
func (p *Parser) parseMultiplication() (Expression, error) {
	left, err := p.parseExponentiation()
	if err != nil {
		return nil, err
	}

	for {
		tok := p.current()
		if tok == nil {
			break
		}

		if tok.Type == spreadsheet.TokenOperator && (tok.Value == "*" || tok.Value == "/") {
			op := tok.Value
			p.advance()
			right, err := p.parseExponentiation()
			if err != nil {
				return nil, err
			}
			left = &BinaryOp{Left: left, Right: right, Operator: op}
		} else {
			break
		}
	}

	return left, nil
}

// parseExponentiation parses exponentiation (^).
// This is right-associative: 2^3^4 = 2^(3^4).
func (p *Parser) parseExponentiation() (Expression, error) {
	left, err := p.parseUnary()
	if err != nil {
		return nil, err
	}

	tok := p.current()
	if tok != nil && tok.Type == spreadsheet.TokenOperator && tok.Value == "^" {
		p.advance()
		// Right-associative: recursively parse the right side
		right, err := p.parseExponentiation()
		if err != nil {
			return nil, err
		}

		return &BinaryOp{Left: left, Right: right, Operator: "^"}, nil
	}

	return left, nil
}

// parseUnary parses unary operators (-, +).
func (p *Parser) parseUnary() (Expression, error) {
	tok := p.current()
	if tok == nil {
		return nil, fmt.Errorf("unexpected end of formula")
	}

	if tok.Type == spreadsheet.TokenOperator && (tok.Value == "-" || tok.Value == "+") {
		op := tok.Value
		p.advance()
		operand, err := p.parseUnary()
		if err != nil {
			return nil, err
		}

		return &UnaryOp{Operand: operand, Operator: op}, nil
	}

	return p.parsePrimary()
}

// parsePrimary parses primary expressions (literals, references, functions, parentheses).
//
//nolint:revive // cyclomatic: parser needs many branches
func (p *Parser) parsePrimary() (Expression, error) {
	tok := p.current()
	if tok == nil {
		return nil, fmt.Errorf("unexpected end of formula")
	}

	switch tok.Type {
	case spreadsheet.TokenLiteral:
		p.advance()

		return p.parseLiteral(tok.Value)

	case spreadsheet.TokenReference:
		p.advance()

		return p.parseReference(tok.Value)

	case spreadsheet.TokenFunction:
		return p.parseFunction()

	case spreadsheet.TokenParen:
		if tok.Value == "(" {
			p.advance()
			expr, err := p.parseComparison()
			if err != nil {
				return nil, err
			}
			if err := p.expect(spreadsheet.TokenParen, ")"); err != nil {
				return nil, err
			}
			p.advance()

			return expr, nil
		}

		return nil, fmt.Errorf("unexpected token: %s", tok.Value)

	default:
		return nil, fmt.Errorf("unexpected token type: %v", tok.Type)
	}
}

// parseLiteral parses a literal value.
func (p *Parser) parseLiteral(s string) (Expression, error) {
	// Try parsing as number
	if num, err := strconv.ParseFloat(s, 64); err == nil {
		return &Literal{Val: NewNumberValue(num)}, nil
	}

	// Try parsing as boolean
	upper := strings.ToUpper(s)
	if upper == "TRUE" {
		return &Literal{Val: NewBooleanValue(true)}, nil
	}
	if upper == "FALSE" {
		return &Literal{Val: NewBooleanValue(false)}, nil
	}

	// Check for error values
	if strings.HasPrefix(s, "#") {
		return &Literal{Val: NewErrorValue(s)}, nil
	}

	// String literal (remove quotes)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		// Unescape doubled quotes
		unescaped := strings.ReplaceAll(s[1:len(s)-1], `""`, `"`)

		return &Literal{Val: NewStringValue(unescaped)}, nil
	}

	// Treat as string
	return &Literal{Val: NewStringValue(s)}, nil
}

// parseReference parses a cell or range reference.
func (p *Parser) parseReference(s string) (Expression, error) {
	// Check if it's a range reference (contains ':')
	if strings.Contains(s, ":") {
		rangeRef, err := spreadsheet.ParseRangeRef(s)
		if err != nil {
			return nil, fmt.Errorf("invalid range reference %q: %w", s, err)
		}

		return &RangeReference{Range: rangeRef}, nil
	}

	// Parse as cell reference
	cellRef, err := spreadsheet.ParseCellRef(s)
	if err != nil {
		return nil, fmt.Errorf("invalid cell reference %q: %w", s, err)
	}

	return &CellReference{Ref: cellRef}, nil
}

// parseFunction parses a function call.
func (p *Parser) parseFunction() (Expression, error) {
	tok := p.current()
	if tok == nil || tok.Type != spreadsheet.TokenFunction {
		return nil, fmt.Errorf("expected function name")
	}

	funcName := tok.Value
	p.advance()

	// Expect opening parenthesis
	if err := p.expect(spreadsheet.TokenParen, "("); err != nil {
		return nil, err
	}
	p.advance()

	// Parse arguments
	var args []Expression

	// Check for empty argument list
	if tok := p.current(); tok != nil && tok.Type == spreadsheet.TokenParen && tok.Value == ")" {
		p.advance()

		return &FunctionCall{Name: funcName, Args: args}, nil
	}

	// Parse first argument
	arg, err := p.parseComparison()
	if err != nil {
		return nil, err
	}
	args = append(args, arg)

	// Parse remaining arguments
	for {
		tok := p.current()
		if tok == nil {
			return nil, fmt.Errorf("unexpected end of formula in function call")
		}

		if tok.Type == spreadsheet.TokenParen && tok.Value == ")" {
			p.advance()

			break
		}

		if tok.Type == spreadsheet.TokenComma {
			p.advance()
			arg, err := p.parseComparison()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		} else {
			return nil, fmt.Errorf("expected ',' or ')' in function call, got %s", tok.Value)
		}
	}

	return &FunctionCall{Name: funcName, Args: args}, nil
}
