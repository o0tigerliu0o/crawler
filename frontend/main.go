package main

import (
	"crawler/frontend/controller"
	"net/http"
)

func main() {

	//不加这个会导致html中引用到的css不会handle
	http.Handle("/", http.FileServer(
		http.Dir("/Users/liuchang/go/gopath/src/crawler/frontend/view/")))
	/*	handler := controller.SearchResultHandler{}
		http.HandleFunc("/search", handler.ServeHTTP)*/

	http.Handle("/search", controller.CreateSearchResultHandler("/Users/liuchang/go/gopath/src/crawler/frontend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
