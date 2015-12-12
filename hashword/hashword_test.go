package hashword

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	line, err := Search("apple")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(line)
}

func TestLookUp(t *testing.T) {
	food, err := LookUp("hihao")
	if err != nil {
		t.Error(err)
	}
	food.PrintAll()
}
