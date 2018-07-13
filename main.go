package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nemith/goline"
	"github.com/yedamao/dict/spider"
)

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr,
		"usage: dict \n"+
			"       dict word\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func autoCompletHandler(l *goline.GoLine) (bool, error) {
	fmt.Println("I can't do that")
	return false, nil
}

func circle() {
	gl := goline.NewGoLine(goline.StringPrompt("dict:>"))
	const TAB rune = 9
	gl.AddHandler(TAB, autoCompletHandler)
	for {
		data, err := gl.Line()
		if err != nil {
			if err == goline.UserTerminatedError {
				fmt.Println("\nBye!")
				return
			} else {
				panic(err)
			}
		}
		fmt.Println()

		find(data)
	}
}

func find(word string) {
	food := spider.Spider(word)
	food.PrintAll()
}

func main() {

	// circle mode
	if len(os.Args) == 1 {
		circle()
		return
	}

	// once mode
	word := os.Args[1]
	find(word)
}
