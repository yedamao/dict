package server

import (
    "fmt"
    // "io/ioutil"
    "net/http"
    "html/template"

    spider "github.com/logindaveye/dict/spider"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("/home/dave/newdict/src/server/first.html")
    t.Execute(w, nil)
}

func resultPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method", r.Method)

    word := r.FormValue("word")
    food := spider.Spider(word)
    fmt.Fprintf(w, "\n" + food.Word + "\n" + food.Pronounce + "\n" + food.Meaning)
}

func Web(port string) {
    http.HandleFunc("/", indexPage)
    http.HandleFunc("/resultPage", resultPage)
    http.ListenAndServe(port, nil)
}
