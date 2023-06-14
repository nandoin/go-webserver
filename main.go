package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("argo automatic k8s deployment kustomized"))
	})
	http.ListenAndServe(":8090", nil)
}
