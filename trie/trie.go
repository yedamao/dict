package trie

import (
	"container/list"
)

const R = 26

func charIndex(c byte) int {
	return int(c - 'a')
}

type node struct {
	val  int
	next [R]*node
}

type Trie struct {
	root *node
}

func New() *Trie {
	return &Trie{}
}

func (t *Trie) Put(key string, val int) {
	t.root = put(t.root, []byte(key), val, 0)
}

func put(x *node, key []byte, val, d int) *node {
	if x == nil {
		x = &node{}
	}

	if d == len(key) {
		x.val = val
		return x
	}

	c := key[d]
	x.next[charIndex(c)] = put(x.next[charIndex(c)], key, val, d+1)
	return x
}

func (t *Trie) Get(key string) int {
	x := get(t.root, []byte(key), 0)
	if x == nil {
		return 0
	}

	return x.val
}

func get(x *node, key []byte, d int) *node {
	if x == nil {
		return nil
	}

	if d == len(key) {
		return x
	}

	c := key[d]
	return get(x.next[charIndex(c)], key, d+1)
}

func (t *Trie) KeyWithPrefix(pre string) *list.List {
	q := list.New()
	collect(get(t.root, []byte(pre), 0), []byte(pre), q)
	return q
}

func collect(x *node, pre []byte, q *list.List) {
	if x == nil {
		return
	}

	if x.val != 0 {
		q.PushFront(string(pre))
	}

	for i := 0; i < R; i++ {
		collect(x.next[i], append(pre, byte(i+'a')), q)
	}
}
