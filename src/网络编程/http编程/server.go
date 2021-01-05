package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go"))
}
func main() {
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
}
