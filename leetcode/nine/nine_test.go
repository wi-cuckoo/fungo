package nine

import "testing"

func TestPrisonAfterNDays(t *testing.T) {
	cells := []int{0, 1, 0, 1, 1, 0, 0, 1}
	prisonAfterNDays(cells, 7)
	for i, v := range []int{0, 0, 1, 1, 0, 0, 0, 0} {
		if cells[i] != v {
			t.Errorf("cells[%d] should be %d, but got %d", i, v, cells[i])
		}
	}

	cells = []int{1, 0, 0, 1, 0, 0, 1, 0}
	prisonAfterNDays(cells, 1000000000)
	for i, v := range []int{0, 0, 1, 1, 1, 1, 1, 0} {
		if cells[i] != v {
			t.Errorf("cells[%d] should be %d, but got %d", i, v, cells[i])
		}
	}
}
