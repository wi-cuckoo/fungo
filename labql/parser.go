package labql

import (
	"fmt"
	"io"
	"strings"
)

// Parser represents an LabQL parser.
type Parser struct {
	s      *BufScanner
	params map[string]interface{}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewBufScanner(r)}
}

// Unscan pushes the previously read token back onto the buffer.
func (p *Parser) Unscan() { p.s.Unscan() }

// Scan returns the next token from the underlying scanner.
func (p *Parser) Scan() (tok Token, pos Pos, lit string) { return p.s.Scan() }

// ScanIgnoreWhitespace scans the next non-whitespace and non-comment token.
func (p *Parser) ScanIgnoreWhitespace() (Token, Pos, string) {
	for {
		tok, pos, lit := p.Scan()
		if tok == WS || tok == COMMENT {
			continue
		}
		return tok, pos, lit
	}
}

// ParseExpr parses an expression.
func (p *Parser) ParseExpr() (Expr, error) {
	var err error
	// Dummy root node.
	root := &BinaryExpr{}

	// Parse a non-binary expression type to start.
	// This variable will always be the root of the expression tree.
	root.RHS, err = p.parseUnaryExpr()
	if err != nil {
		return nil, err
	}

	// Loop over operations and unary exprs and build a tree based on precendence.
	for {
		// If the next token is NOT an operator then return the expression.
		op, _, _ := p.ScanIgnoreWhitespace()
		if !op.isOperator() {
			p.Unscan()
			return root.RHS, nil
		}

		// Otherwise parse the next expression.
		var rhs Expr
		if rhs, err = p.parseUnaryExpr(); err != nil {
			return nil, err
		}

		// Find the right spot in the tree to add the new expression by
		// descending the RHS of the expression tree until we reach the last
		// BinaryExpr or a BinaryExpr whose RHS has an operator with
		// precedence >= the operator being added.
		for node := root; ; {
			r, ok := node.RHS.(*BinaryExpr)
			if !ok || r.Op.Precedence() >= op.Precedence() {
				// Add the new expression here and break.
				node.RHS = &BinaryExpr{LHS: node.RHS, RHS: rhs, Op: op}
				break
			}
			node = r
		}
	}
}

// parseUnaryExpr parses an non-binary expression.
func (p *Parser) parseUnaryExpr() (Expr, error) {
	// If the first token is a LPAREN then parse it as its own grouped expression.
	if tok, _, _ := p.ScanIgnoreWhitespace(); tok == LPAREN {
		expr, err := p.ParseExpr()
		if err != nil {
			return nil, err
		}

		// Expect an RPAREN at the end.
		if tok, pos, _ := p.ScanIgnoreWhitespace(); tok != RPAREN {
			return nil, fmt.Errorf("found %s expect ), at line %d, char %d", tok, pos.Line+1, pos.Char+1)
		}

		return &ParenExpr{Expr: expr}, nil
	}
	p.Unscan()

	// Read next token.
	tok, pos, lit := p.ScanIgnoreWhitespace()
	switch tok {
	case STRING:
		return &StringLiteral{Val: lit}, nil
	default:
		return nil, fmt.Errorf("found %s, at line %d, char %d", tok, pos.Line+1, pos.Char+1)
	}
}

// ParseExpr parses an expression string and returns its AST representation.
func ParseExpr(s string) (Expr, error) { return NewParser(strings.NewReader(s)).ParseExpr() }
