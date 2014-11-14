package rw

import (
    "os"
    "fmt"
)

func Read(path string) []byte {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Failed to open:", path)
        os.Exit(1)
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        fmt.Println("Get stat error")
        os.Exit(2)
    }

    data := make([]byte, stat.Size())
    _, err = file.Read(data)
    if err != nil {
        fmt.Println("Can't read from:", path)
    }

    return data
}

func WriteLine(line, path string) error{
    file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Println("Failed to Create output file", path)
        return err
    }
    defer file.Close()

    file.WriteString(line + "\n")

    return nil
}

func Write(content , path string) error{
    file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Println("Failed to open save word file")
        return err
    }
    defer file.Close()

    file.WriteString(content)

    return nil
}
