package spider

import (
    "regexp"
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
)

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func findWord(content string) {
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
}

func addFoods(food [][]string) string {
    wordLine := ""
    for i := 0; i < len(food); i++ {
        wordLine += food[i][1]
    }

    fmt.Println(wordLine)
    return wordLine
}


func Spider(word string) {
    URL := "http://dict.youdao.com/search?q=" + word + "&keyfrom=dict.index"
    // URL := "http://dict.youdao.com/search?q=go&keyfrom=dict.index"
    res, err := http.Get(URL)
    checkError(err)

    robots, err := ioutil.ReadAll(res.Body)
    checkError(err)
    res.Body.Close()

    findWord(string(robots))
}
