package labql

// Token is a lexical token of the LabQL language.
type Token int

// These are a comprehensive list of InfluxQL language tokens.
const (
	// ILLEGAL Token, EOF, WS are Special InfluxQL tokens.
	ILLEGAL Token = iota
	EOF
	WS
	COMMENT

	STRING

	operatorBeg
	AND // AND
	OR  // OR
	EQ  // =
	NEQ // !=
	IN  // IN
	operatorEnd

	LPAREN // (
	RPAREN // )
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",

	STRING: "STRING",

	AND:    "AND",
	OR:     "OR",
	EQ:     "=",
	NEQ:    "!=",
	IN:     "IN",
	LPAREN: "(",
	RPAREN: ")",
}

var keywords = map[string]Token{
	"AND": AND,
	"OR":  OR,
	"IN":  IN,
}

// String returns the string representation of the token.
func (t Token) String() string {
	if t >= 0 && t < Token(len(tokens)) {
		return tokens[t]
	}
	return ""
}

// Precedence returns the operator precedence of the binary operator token.
func (t Token) Precedence() int {
	switch t {
	case OR:
		return 1
	case AND:
		return 2
	case EQ, NEQ, IN:
		return 3
	}
	return 0
}

// isOperator returns true for operator tokens.
func (t Token) isOperator() bool { return t > operatorBeg && t < operatorEnd }

// Pos specifies the line and character position of a token.
// The Char and Line are both zero-based indexes.
type Pos struct {
	Line int
	Char int
}
