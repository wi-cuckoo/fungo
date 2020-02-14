package labql_test

import (
	"fmt"
	"testing"

	"github.com/wi-cuckoo/fungo/labql"
)

func TestParseExpr(t *testing.T) {
	s := "env = prod AND os = unix"
	e, err := labql.ParseExpr(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(e.String())
}
