package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
    spider "github.com/logindaveye/dict/spider"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func readDictionary() []byte {
	//open dictionary file
	file, err := os.Open("dictionary")
	checkError(err)
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	checkError(err)

	//read the file
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	checkError(err)

	return data
}

func searchWordLine(word string) string {
	// binary search
	data := string(readDictionary())
	wordLines := strings.Split(data, "\n")
	low := 0
	high := len(wordLines)
	index := -1

	for low <= high {
		h := low + (high-low)/2
		target := strings.Split(wordLines[h], ":")[0]
		// if strings.ToLower(target) == strings.ToLower(word) {
		if target == word {
			index = h
			break
		} else if strings.ToLower(target) < strings.ToLower(word) {
			low = h + 1
		} else {
			high = h - 1
		}
	}

	if index == -1 {
        spider.Spider(word)
        return ""
	} else {
		return wordLines[index]
	}
}

func printWordLine(wordLine string) {
	for i := 0; i < len(strings.Split(wordLine, ":")); i++ {
		fmt.Println(strings.Split(wordLine, ":")[i])
	}
	fmt.Println("_______________________________")
}

func main() {
    if len(os.Args[:]) > 1 {
        args := os.Args[1]
        if args == "--help" {
            fmt.Println("help")
        } else {
            spider.Spider(args)
        }
    } else {
        reader := bufio.NewReader(os.Stdin)
        
        fmt.Print("Open Source Dictionary\n")
        
        for {
            fmt.Print("(OSD1.1):>")
            line, _ := reader.ReadString('\n')
            word := strings.Split(line, "\n")[0]
            wordLine := searchWordLine(word)
            if wordLine != "" {
                printWordLine(wordLine)
            } else {
                fmt.Println("Oooooooo!")
            }
        }

    }
}
