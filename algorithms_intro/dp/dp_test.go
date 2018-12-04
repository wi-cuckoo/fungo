package dp

import (
	"testing"
)

func TestCoinInLine(t *testing.T) {
	if CoinInLine(3) != false {
		t.Error("when 3 coins, first should win")
	}
}
