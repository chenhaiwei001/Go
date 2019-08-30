package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>测试一下</h1>")
}


func main() {
	server := &http.Server{
		Addr: "0.0.0.0:80",
	}
	http.HandleFunc("/hello", hello)
	server.ListenAndServe()
}