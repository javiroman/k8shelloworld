package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		log.Printf("%s - [%s %s %s] 200 [%s]",
			r.RemoteAddr, r.Method, r.URL.Path, r.Proto, userAgent)
		w.Header().Set("Server", "Hello/v2")
		fmt.Fprintf(w, "CPUs: %d\n", runtime.NumCPU())
		fmt.Fprintf(w, "Hostname: %s\n", hostname)
	})

	log.Println("Starting Hello service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
