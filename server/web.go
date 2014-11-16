package server

import (
	"fmt"
	// "io/ioutil"
	"html/template"
	"net/http"

	lookup "github.com/logindaveye/dict/lookup"
	spider "github.com/logindaveye/dict/spider"
)

var tmplPath string = "/home/dave/gocode/src/github.com/logindaveye/dict/server/tmpl/"

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(tmplPath + "index.html")
	t.Execute(w, nil)
}

func resultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)

	word := r.FormValue("word")

	food := lookup.Lookup(word)
	if food.Word == "" {
		food = spider.Spider(word)
	}

	fmt.Fprintf(w, "\n"+food.Word+"\n"+food.Pronounce+"\n"+food.Meaning)
}

func Web(port string) {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/resultPage", resultPage)
	http.ListenAndServe(port, nil)
}
