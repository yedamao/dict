package lookup

import (
    "strings"
    spider "github.com/logindaveye/dict/spider"
    rw "github.com/logindaveye/dict/rw"
)

// const DICTIONARY_PATH string = "/usr/share/dict/dictionary"
const DICTIONARY_PATH string = "/usr/share/dict/spider_word"

func searchWordLine(word string) string {
    //read dictionary file
	data := rw.Read(DICTIONARY_PATH)
	wordLines := strings.Split(string(data), "\n")

	// binary search
	low := 0
	high := len(wordLines)
	index := -1

	for low <= high { //fix bug low and high can't equal
		h := (low + high)/2
        if h >= len(wordLines) { //if index out of range return -1
            index = -1
            break
        }
		target := strings.Split(wordLines[h], "#")[0]
		// if strings.ToLower(target) == strings.ToLower(word) {
		if target == word {
			index = h
			break
		// } else if strings.ToLower(target) < strings.ToLower(word) {
		} else if target < word {
			low = h + 1
		} else {
			high = h - 1
		}
	}

	if index == -1 {
		return "not found"
    } else {
		return wordLines[index]
	}
}

func Lookup(word string) spider.Foods{
    food := new(spider.Foods)

    wordLine := searchWordLine(word)
    if wordLine == "not found" {
        return *food
    }

    food.Word = strings.Split(wordLine, "#")[0]
    food.Pronounce = strings.Split(wordLine, "#")[1]
    food.Meaning = strings.Split(wordLine, "#")[2]

    return *food
}
