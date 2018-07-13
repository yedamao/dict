package main

import (
	hashword "github.com/logindave/dict/hashword"
	spider "github.com/logindave/dict/spider"

	goline "github.com/nemith/goline"

	"flag"
	"fmt"
	"os"
)

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

func usage() {
	fmt.Fprintf(os.Stderr,
		"usage: dict \n"+
			"       dict word\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func find(word string) {
	food, err := hashword.LookUp(word)
	if err != nil {
		food = spider.Spider(word)
		food.WriteAll("/home/dave/.dictionary")
		fmt.Println("from internet")
	}
	food.PrintAll()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) == 1 {
		circle()
	} else if os.Args[1][0] != '-' {
		word := os.Args[1]
		find(word)
	} else {
	}
}
