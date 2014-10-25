package spider

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func findWord(content string) string {
	wordLine := ""

	r_word, _ := regexp.Compile("<span class=\"keyword\">([a-zA-Z]*)</span>")

	word := r_word.FindAllStringSubmatch(string(content), -1)
	wordLine += addFoods(word)

	r_pronounce, _ := regexp.Compile("<span class=\"phonetic\">(.*)</span>")

	wordLine += addFoods(r_pronounce.FindAllStringSubmatch(string(content), 2))

	r_transContainer, _ := regexp.Compile("<div class=\"trans-container\">\\s*<ul>\\s*(<li>.*</li>\\s*)*\\s*</ul>")
	meaning := r_transContainer.FindString(string(content))

	r_meaning, _ := regexp.Compile("<li>(.*)</li>")
	wordLine += addFoods(r_meaning.FindAllStringSubmatch(string(meaning), -1))

	return wordLine
}

func addFoods(food [][]string) string {
	wordLine := ""
	for i := 0; i < len(food); i++ {
		wordLine += food[i][1]
	}

	return wordLine + ":"
}

func Spider(word string) string {
	//return result
	URL := "http://dict.youdao.com/search?q=" + word + "&keyfrom=dict.index"
	res, err := http.Get(URL)
	checkError(err)

	robots, err := ioutil.ReadAll(res.Body)
	checkError(err)
	res.Body.Close()

	return findWord(string(robots))
}
