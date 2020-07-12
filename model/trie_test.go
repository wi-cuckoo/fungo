package model

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tr := &Trie{next: [26]*Trie{}}
	tr.Insert("louis")
	tr.Insert("wei")
	tr.Insert("longtimetosee")
	t.Log(tr.Exist("a"))
}
