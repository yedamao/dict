package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peterh/liner"
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

func loop() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	for {
		if word, err := line.Prompt("dict:>"); err == nil {
			find(word)
			line.AppendHistory(word)
		} else if err == liner.ErrPromptAborted {
			fmt.Println("Bye!")
			return
		} else {
			panic(err)
		}
	}
}

func find(word string) {
	food := spider.Spider(word)
	food.PrintAll()
}

func main() {
	// loop mode
	if len(os.Args) == 1 {
		loop()
		return
	}

	// once mode
	word := os.Args[1]
	find(word)
}
