package labql

import "fmt"

// BinaryExpr represents an operation between two expressions.
type BinaryExpr struct {
	Op  Token
	LHS Expr
	RHS Expr
}

func (*BinaryExpr) node() {}
func (*BinaryExpr) expr() {}

// String returns a string representation of the binary expression.
func (e *BinaryExpr) String() string {
	switch e.Op {
	case AND, OR:
		return fmt.Sprintf("%s %s %s", e.LHS, e.Op, e.RHS)
	default:
		return fmt.Sprintf("(key = %s AND value %s %s)", e.LHS, e.Op, e.RHS)
	}
}

// ParenExpr represents a parenthesized expression.
type ParenExpr struct {
	Expr Expr
}

// String returns a string representation of the parenthesized expression.
func (e *ParenExpr) String() string { return fmt.Sprintf("(%s)", e.Expr.String()) }
func (*ParenExpr) node()            {}
func (*ParenExpr) expr()            {}

// StringLiteral represents a string literal.
type StringLiteral struct {
	Val string
}

// String returns a string representation of the literal.
func (l *StringLiteral) String() string { return l.Val }
func (*StringLiteral) node()            {}
func (*StringLiteral) expr()            {}

// Expr represents an expression that can be evaluated to a value.
type Expr interface {
	Node
	// expr is unexported to ensure implementations of Expr
	// can only originate in this package.
	expr()
}

// Node represents a node in the InfluxDB abstract syntax tree.
type Node interface {
	// node is unexported to ensure implementations of Node
	// can only originate in this package.
	node()
	String() string
}
