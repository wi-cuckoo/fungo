package model

// Trie 中文名叫字典树来着
type Trie struct {
	next [26]*Trie //假设只有小写字母
	end  bool
}

// Insert a word
func (t *Trie) Insert(s string) {
	cur := t
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if cur.next[idx] == nil {
			cur.next[idx] = &Trie{next: [26]*Trie{}}
		}
		cur = cur.next[idx]
	}
	cur.end = true
}

// Exist query
func (t *Trie) Exist(s string) bool {
	cur := t
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}
	return cur.end
}
