package main

import (
    "fmt"
    "os"
    "log"
)

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}
func main() {
    //open dictionary file
    file , err := os.Open("dictionary")
    checkError(err)
    data := make([]byte, 100)
    count, err := file.Read(data)
    checkError(err)
    
    fmt.Println("read %d bytes: %q\n", count, data[:count])
}
