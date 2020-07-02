package labql_test

import (
	"testing"

	"github.com/wi-cuckoo/fungo/labql"
)

func TestParseExpr(t *testing.T) {
	s := "env= prod AND (os != unix OR dept=wsd)"
	expr, err := labql.ParseExpr(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(expr)
}
