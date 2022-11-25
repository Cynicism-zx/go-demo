package main

import (
	"fmt"
	"net/http"

	"go-demo/base/http/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.BodyLimit(hello))

	mux.HandleFunc("/admin", middleware.Auth(hello))

	go fmt.Println("server staring...")

	if err := http.ListenAndServe(":8081", middleware.IPRateLimit(mux)); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
