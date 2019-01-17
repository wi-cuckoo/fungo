package seven

import "testing"

func TestRotateString(t *testing.T) {
	a, b := "gcmbf", "fgcmb"
	if !rotateString(a, b) {
		t.Error("should be true")
	}
}
