package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		version := os.Getenv("VERSION")
		if version == "" {
			version = "1.0.0" // 默认版本号
		}
		fmt.Fprintf(w, "<h1>Hello, CICD Master!</h1><p>Welcome to our Go web app.</p><br><strong>Version: %s</strong>", version)
	})

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
