package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPutGet(t *testing.T) {
	testcases := []struct {
		key string
		val int
	}{
		{"a", 1},
		{"ape", 2},
		{"apple", 3},
	}

	trie := New()
	for _, tc := range testcases {
		trie.Put(tc.key, tc.val)
	}

	for _, tc := range testcases {
		require.Equal(t, tc.val, trie.Get(tc.key))
	}

	require.Equal(t, 3, trie.KeyWithPrefix("a").Len())
	require.Equal(t, 2, trie.KeyWithPrefix("ap").Len())
}
