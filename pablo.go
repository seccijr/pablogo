package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})

	var port = os.Getenv("PORT")
	if port == "" {
		panic("No port environment variable found")
	}
	var addr = fmt.Sprintf("%s:%s", os.Getenv("HOST"), port)
	http.ListenAndServe(addr, nil)
}
