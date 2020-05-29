package main

import (
	"fmt"
	"net/http"
)

func Sayhello(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	http.HandleFunc("/", Sayhello)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Http Service start failed!")
		return
	}

}
