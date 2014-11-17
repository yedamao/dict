package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	// rw "github.com/logindaveye/dict/rw"
	"github.com/logindave/dict/rw"
)

type Foods struct {
	Word      string
	Pronounce string
	Meaning   string
}

func (food Foods) PrintAll() {
	fmt.Println(food.Word)
	for _, x := range strings.Split(food.Pronounce, "|") {
		fmt.Println(x)
	}
	for _, x := range strings.Split(food.Meaning, "|") {
		fmt.Println(x)
	}
}

func (food Foods) WriteAll(path string) {
	//: filt word pronouce meaning
	// if food.Word != "" && food.Pronounce != "" && food.Meaning != "" {
	if food.Word != "" && food.Meaning != "" { //modify pronounce is nil also can save
		wordline := food.Word + "#" + food.Pronounce + "#" + food.Meaning
		err := rw.WriteLine(wordline, path)
		if err != nil {
			fmt.Println("something err")
		} else {
			fmt.Println(food.Word, "write sucessful")
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func Spider(word string) Foods {
	URL := "http://dict.youdao.com/search?q=" + word + "&keyfrom=dict.index"

	resp, err := http.Get(URL)
	if err != nil {
		return *new(Foods)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return findFoods(string(body))
}

func findFoods(body string) Foods {
	food := new(Foods)
	r_word, _ := regexp.Compile("<span class=\"keyword\">([a-zA-Z]*)</span>")
	if r_word.FindAllStringSubmatch(body, -1) == nil {
		food.Word = ""
	} else {
		food.Word = r_word.FindAllStringSubmatch(body, -1)[0][1]
	}

	r_pronounce, _ := regexp.Compile("<span class=\"phonetic\">(.*)</span>")
	food.Pronounce = addFoods(r_pronounce.FindAllStringSubmatch(body, -1))

	r_transContainer, _ := regexp.Compile("<div class=\"trans-container\">\\s*<ul>\\s*(<li>.*</li>\\s*)*\\s*</ul>")
	transContainer := r_transContainer.FindString(body)

	r_meaning, _ := regexp.Compile("<li>(.*)</li>")
	food.Meaning = addFoods(r_meaning.FindAllStringSubmatch(transContainer, -1))

	return *food
}

func addFoods(rawFoods [][]string) string {
	foods := ""
	for _, x := range rawFoods {
		//| is filter
		foods += x[1] + "|"
	}
	if foods != "" {
		foods = foods[:len(foods)-1]
	}

	return foods
}
