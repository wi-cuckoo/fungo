package four

import "testing"

func TestRemoveKdigits(t *testing.T) {
	t.Log(removeKdigits("1234219", 5))
	t.Log(removeKdigits("100200", 1))
	t.Log(removeKdigits("10", 1))
}
