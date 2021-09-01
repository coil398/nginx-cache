package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/ramen", ramen)
	mux.HandleFunc("/nginx", nginxRamen)

	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func ramen(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/assets/ramen.jpg")
}

func nginxRamen(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("nginxRamen")
	w.Header().Set("X-Accel-Redirect", "/assets/ramen.jpg")
	fmt.Fprintln(w, "nginx-ramen")
}
