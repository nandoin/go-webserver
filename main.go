package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("subiu rapaziada"))
	})
	http.ListenAndServe(":8090", nil)
}