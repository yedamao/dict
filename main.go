package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/peterh/liner"
	"github.com/yedamao/dict/spider"
	"github.com/yedamao/dict/trie"
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

const SPEAKER_CMD = "say"

func checkSpeaker() error {
	if _, err := exec.LookPath(SPEAKER_CMD); err != nil {
		return err
	}

	return exec.Command(SPEAKER_CMD, "welcome").Run()
}

func say(word string) {
	exec.Command(SPEAKER_CMD, word).Run()
}

func loop() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	trie, err := initCompleter("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}

	line.SetCompleter(func(line string) (c []string) {
		if len(line) < 3 {
			return
		}
		q := trie.KeyWithPrefix(line)
		for q.Len() > 0 {
			c = append(c, q.Remove(q.Back()).(string))
		}
		return
	})

	isSupportSpeaker := checkSpeaker() == nil

	for {
		if word, err := line.Prompt("dict:>"); err == nil {
			food := spider.Spider(word)
			food.PrintAll()
			if isSupportSpeaker {
				say(word)
			}
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

func initCompleter(file string) (*trie.Trie, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	trie := trie.New()

	reader := bufio.NewReader(f)
	for {
		word, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		word = strings.TrimSpace(word)
		if isLowercaseAlphabet(word) {
			trie.Put(word, 1)
		}
	}
	return trie, nil
}

func isLowercaseAlphabet(s string) bool {
	for _, b := range s {
		if b < 'a' || b > 'z' {
			return false
		}
	}
	return true
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
