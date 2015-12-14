package autocomplete

import (
	"fmt"
	"github.com/logindave/dict/rw"
	"strings"
	"testing"
)

var words []string

func init() {
	data := rw.Read("./words")
	words = strings.Split(string(data), "\n")
}

func TestInsert(t *testing.T) {
	trie := NewTrie()

	for _, val := range words {
		trie.Insert(val)
	}
}

func TestAutoComplete(t *testing.T) {
	trie := NewTrie()
	for _, val := range words {
		trie.Insert(val)
	}

	results, err := trie.AutoComplete("apple")
	if err != nil {
		t.Error("An error ocurred: " + err.Error())
	}
	fmt.Println(results)
}
