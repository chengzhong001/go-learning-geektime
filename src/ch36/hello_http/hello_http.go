package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello world")
	})
	http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
		t:= time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		rw.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
}
