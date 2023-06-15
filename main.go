package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("MESSAGE")))
	})
	http.ListenAndServe(":8090", nil)
}
