// Package spreadsheet provides SpreadsheetML support for Excel documents.
//
//nolint:revive // file-length-limit: formula rewriting logic is cohesive
package spreadsheet

import (
	"errors"
	"strings"
	"unicode"
)

// TokenType represents the type of formula token.
type TokenType int

const (
	// TokenOperator represents operators: +, -, *, /, =, <>, <=, >=, &, ^.
	TokenOperator TokenType = iota
	// TokenFunction represents function names (followed by '(').
	TokenFunction
	// TokenReference represents cell or range references (A1, $A$1, Sheet1!A1, A1:D10).
	TokenReference
	// TokenLiteral represents string, numeric, or boolean literals.
	TokenLiteral
	// TokenParen represents parentheses: (, ).
	TokenParen
	// TokenComma represents comma separator.
	TokenComma
	// TokenColon represents colon (part of range, but sometimes standalone).
	TokenColon
)

// Token represents a single token in a formula.
type Token struct {
	Type  TokenType
	Value string
}

// ErrInvalidFormula is returned when formula tokenization fails.
var ErrInvalidFormula = errors.New(
	"invalid formula",
)

// tokenizeFormula splits a formula into tokens.
// The formula should include the leading '=' if it's a cell formula.
//
//nolint:revive // function-length, cognitive-complexity: tokenization is complex
func tokenizeFormula(
	formula string,
) ([]Token, error) {
	if formula == "" {
		return nil, ErrInvalidFormula
	}

	var tokens []Token
	i := 0

	// Skip leading '=' if present
	if formula[0] == '=' {
		i = 1
	}

	for i < len(formula) {
		// Skip whitespace
		if unicode.IsSpace(rune(formula[i])) {
			i++

			continue
		}

		ch := formula[i]

		// Handle string literals
		if ch == '"' {
			start := i
			i++
			for i < len(formula) {
				if formula[i] == '"' {
					// Check for escaped quote
					if i+1 < len(formula) &&
						formula[i+1] == '"' {
						i += 2

						continue
					}
					i++

					break
				}
				i++
			}
			tokens = append(tokens, Token{
				Type:  TokenLiteral,
				Value: formula[start:i],
			})

			continue
		}

		// Handle operators
		if ch == '+' || ch == '-' || ch == '*' ||
			ch == '/' ||
			ch == '^' ||
			ch == '&' {
			tokens = append(tokens, Token{
				Type:  TokenOperator,
				Value: string(ch),
			})
			i++

			continue
		}

		// Handle two-character operators
		if ch == '=' || ch == '<' || ch == '>' {
			if i+1 < len(formula) {
				next := formula[i+1]
				if (ch == '<' && next == '>') ||
					(ch == '<' && next == '=') ||
					(ch == '>' && next == '=') {
					tokens = append(tokens, Token{
						Type:  TokenOperator,
						Value: formula[i : i+2],
					})
					i += 2

					continue
				}
			}
			tokens = append(tokens, Token{
				Type:  TokenOperator,
				Value: string(ch),
			})
			i++

			continue
		}

		// Handle parentheses
		if ch == '(' || ch == ')' {
			tokens = append(tokens, Token{
				Type:  TokenParen,
				Value: string(ch),
			})
			i++

			continue
		}

		// Handle comma
		if ch == ',' {
			tokens = append(tokens, Token{
				Type:  TokenComma,
				Value: ",",
			})
			i++

			continue
		}

		// Colons are now handled as part of range references in the identifier parsing
		// If we encounter a standalone colon, it's an error
		if ch == ':' {
			return nil, ErrInvalidFormula
		}

		// Handle numeric literals
		if unicode.IsDigit(rune(ch)) ||
			(ch == '.' && i+1 < len(formula) && unicode.IsDigit(rune(formula[i+1]))) {
			start := i
			hasDecimal := ch == '.'
			if hasDecimal {
				i++
			}
			for i < len(formula) {
				c := formula[i]
				switch {
				case unicode.IsDigit(rune(c)):
					i++
				case c == '.' && !hasDecimal:
					hasDecimal = true
					i++
				case c == 'E' || c == 'e':
					// Handle scientific notation
					i++
					if i < len(formula) &&
						(formula[i] == '+' || formula[i] == '-') {
						i++
					}
				default:
					goto endNumeric
				}
			}
		endNumeric:
			tokens = append(tokens, Token{
				Type:  TokenLiteral,
				Value: formula[start:i],
			})

			continue
		}

		// Handle identifiers (function names, cell references, named ranges, TRUE/FALSE)
		if unicode.IsLetter(rune(ch)) ||
			ch == '$' ||
			ch == '\'' ||
			ch == '_' {
			start := i

			// Handle quoted sheet names
			if ch == '\'' {
				i++
				for i < len(formula) && formula[i] != '\'' {
					i++
				}
				if i < len(formula) {
					i++ // Skip closing quote
				}
				if i < len(formula) &&
					formula[i] == '!' {
					i++ // Include the '!'
				}
			}

			// Continue parsing the identifier
			for i < len(formula) {
				c := formula[i]
				if unicode.IsLetter(rune(c)) ||
					unicode.IsDigit(rune(c)) ||
					c == '_' ||
					c == '.' ||
					c == '$' ||
					c == '!' {
					i++
				} else {
					break
				}
			}

			// Check for range reference (cell:cell)
			if i < len(formula) &&
				formula[i] == ':' {
				// This might be a range reference, consume the colon and continue
				i++ // Skip the ':'

				// Continue parsing the second part of the range
				for i < len(formula) {
					c := formula[i]
					if unicode.IsLetter(
						rune(c),
					) ||
						unicode.IsDigit(
							rune(c),
						) ||
						c == '_' ||
						c == '.' ||
						c == '$' {
						i++
					} else {
						break
					}
				}
			}

			value := formula[start:i]

			// Check if this is a function (followed by '(')
			if i < len(formula) {
				// Skip whitespace to check for '('
				j := i
				for j < len(formula) && unicode.IsSpace(rune(formula[j])) {
					j++
				}
				if j < len(formula) &&
					formula[j] == '(' {
					tokens = append(tokens, Token{
						Type:  TokenFunction,
						Value: value,
					})

					continue
				}
			}

			// Check if this looks like a cell reference or range reference
			if isCellOrRangeRef(value) {
				tokens = append(tokens, Token{
					Type:  TokenReference,
					Value: value,
				})
			} else {
				// Treat as literal (TRUE, FALSE, or named range)
				tokens = append(tokens, Token{
					Type:  TokenLiteral,
					Value: value,
				})
			}

			continue
		}

		// Unknown character
		return nil, ErrInvalidFormula
	}

	return tokens, nil
}

// isCellOrRangeRef checks if a string looks like a cell or range reference.
// This is a heuristic check used during tokenization.
func isCellOrRangeRef(s string) bool {
	// Try parsing as cell reference
	_, err := ParseCellRef(s)
	if err == nil {
		return true
	}

	// Try parsing as range reference
	_, err = ParseRangeRef(s)

	return err == nil
}

// reconstructFormula rebuilds a formula from tokens.
// It adds the leading '=' if not present in the first token.
func reconstructFormula(tokens []Token) string {
	if len(tokens) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteByte('=')

	for i, tok := range tokens {
		// Add spacing for readability where appropriate
		if i > 0 {
			prev := tokens[i-1]
			// Add space before operator (except after open paren)
			if tok.Type == TokenOperator &&
				prev.Type != TokenParen {
				sb.WriteByte(' ')
			}
			// Add space after operator (except before close paren or comma)
			if prev.Type == TokenOperator &&
				tok.Type != TokenParen &&
				tok.Type != TokenComma {
				sb.WriteByte(' ')
			}
			// Add space after comma
			if prev.Type == TokenComma {
				sb.WriteByte(' ')
			}
		}

		sb.WriteString(tok.Value)
	}

	return sb.String()
}

// FormulaRewriter rewrites formulas when rows/columns are shifted.
type FormulaRewriter struct {
	updater *CellReferenceUpdater
}

// NewFormulaRewriter creates a new formula rewriter.
func NewFormulaRewriter(
	operation ShiftOperation,
	currentSheet string,
) *FormulaRewriter {
	return &FormulaRewriter{
		updater: NewCellReferenceUpdater(
			operation,
			currentSheet,
		),
	}
}

// Rewrite updates all cell references in a formula based on the shift operation.
// It tokenizes the formula, updates reference tokens, and reconstructs the formula.
// Returns the updated formula, or the original formula if no updates were needed.
// If a reference is deleted, it's replaced with "#REF!".
//
//nolint:revive // cognitive-complexity: formula rewriting requires branches
func (r *FormulaRewriter) Rewrite(
	formula string,
) (string, error) {
	if formula == "" {
		return formula, nil
	}

	// Preserve leading '=' for reconstruction
	hasEquals := formula[0] == '='

	// Tokenize the formula
	tokens, err := tokenizeFormula(formula)
	if err != nil {
		// If tokenization fails, return original formula
		return formula, err
	}

	// Track if any changes were made
	changed := false

	// Update each reference token
	for i := range tokens {
		if tokens[i].Type != TokenReference {
			continue
		}

		// Try parsing as range reference first (ranges contain ':')
		if containsColon(tokens[i].Value) {
			rangeRef, err := ParseRangeRef(
				tokens[i].Value,
			)
			if err != nil {
				// Not a valid range, might be a cell ref - continue
				continue
			}

			// Update the range
			updatedRange, err := r.updater.UpdateRange(
				rangeRef,
			)
			if err == ErrReferenceDeleted {
				// Replace with #REF!
				tokens[i].Value = "#REF!"
				changed = true

				continue
			}

			// Convert back to string
			updatedStr := updatedRange.String()
			if updatedStr != tokens[i].Value {
				tokens[i].Value = updatedStr
				changed = true
			}
		} else {
			// Try parsing as cell reference
			cellRef, err := ParseCellRef(tokens[i].Value)
			if err != nil {
				// Not a valid cell reference, skip
				continue
			}

			// Update the cell reference
			updatedRef, err := r.updater.UpdateReference(cellRef)
			if err == ErrReferenceDeleted {
				// Replace with #REF!
				tokens[i].Value = "#REF!"
				changed = true

				continue
			}

			// Convert back to string
			updatedStr := updatedRef.String()
			if updatedStr != tokens[i].Value {
				tokens[i].Value = updatedStr
				changed = true
			}
		}
	}

	// If no changes, return original formula
	if !changed {
		return formula, nil
	}

	// Reconstruct the formula
	result := reconstructFormula(tokens)

	// If original didn't have '=', remove it from result
	if !hasEquals && result != "" &&
		result[0] == '=' {
		result = result[1:]
	}

	return result, nil
}

// containsColon checks if a string contains a colon character.
func containsColon(s string) bool {
	return strings.Contains(s, ":")
}
