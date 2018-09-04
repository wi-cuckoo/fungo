package sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := [...]int{31, 41, 59, 26, 41, 58}
	InsertSort(arr[:])
	t.Log(arr)
}
