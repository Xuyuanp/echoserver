package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func echo(w http.ResponseWriter, r *http.Request) {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(data))
}

func main() {
	http.HandleFunc("/", echo)
	panic(http.ListenAndServe(":8080", nil))
}
