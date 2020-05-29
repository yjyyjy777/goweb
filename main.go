package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Sayhello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./tpl/hello.tmpl")
	if err != nil {
		fmt.Println("template init failed!")
		return
	}
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed!")
		return
	}

}

func main() {
	http.HandleFunc("/", Sayhello)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Http Service start failed!")
		return
	}

}
