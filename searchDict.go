package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
    "log"
)

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}

func readDictionary() []byte {
    //open dictionary file
    file , err := os.Open("dictionary")
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

func searchWordLine(word string) {
    data := string(readDictionary())
    wordLines := strings.Split(data, "\n")
    fmt.Println(word)
    fmt.Println(strings.Split(wordLines[0], ":")[0])

    i := sort.Search(len(wordLines), func(i int) bool {
        return word <= strings.Split(wordLines[i], ":")[0]
    })
    fmt.Println(i)
}

func main() {
    // fmt.Println(searchWordLine("apple"))
    searchWordLine("a")
}
