package hashword

import (
	"errors"
	"github.com/logindave/dict/rw"
	"github.com/logindave/dict/spider"
	"strings"
)

const N int32 = 1001

type Item struct {
	line string
	next *Item
}

// table have N slot, type *Item
var table [N]*Item

// sum all char as hash code
func hashword(word string) int32 {
	var sum int32 = 0
	for _, x := range strings.TrimSpace(word) {
		sum += x
	}
	return sum
}

func getWord(line string) string {
	return strings.Split(line, "#")[0]
}

// load data
// setup hashtable
func init() {
	data := rw.Read("/usr/share/dict/spider_word")
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		word := getWord(line)

		code := hashword(word)

		p := new(Item)
		p.line = line
		p.next = nil

		t := table[code%N]
		if t == nil {
			table[code%N] = p
		} else {
			for t.next != nil {
				t = t.next
			}
			t.next = p
		}
	}
}

func Search(word string) (line string, err error) {
	code := hashword(word)

	t := table[code%N]
	if t == nil {
		return "", errors.New("slot is nil")
	} else {
		for {
			if strings.Compare(word, getWord(t.line)) == 0 {
				line = t.line
				break
			} else {
				t = t.next
				if t == nil {
					return "", errors.New("tail, not found")
				}
			}
		}
	}
	return line, nil
}

func LookUp(word string) (spider.Foods, error) {
	food := new(spider.Foods)
	line, err := Search(word)
	if err != nil {
		return *food, err
	} else {
		food.Word = getWord(line)
		food.Pronounce = strings.Split(line, "#")[1]
		food.Meaning = strings.Split(line, "#")[2]

		return *food, nil
	}
}
