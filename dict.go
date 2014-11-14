package main

import (
    spider "github.com/logindaveye/dict/spider"
    lookup "github.com/logindaveye/dict/lookup"
    server "github.com/logindaveye/dict/server"

    goline "github.com/nemith/goline"

    "fmt"
    "os"
    "flag"
)


func circle() {
    gl := goline.NewGoLine(goline.StringPrompt("dict:>"))

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

var http string
func init() {
    flag.StringVar(&http, "http", ":6060", "the address of server")
}

func usage() {
    fmt.Fprintf(os.Stderr,
        "usage: dict \n" +
        "       dict word\n")
    flag.PrintDefaults()
    os.Exit(2)
}

func find(word string) {
    food := lookup.Lookup(word)
    if food.Word == "" {
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
        server.Web(http)
    }
}
