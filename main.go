package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	hostname, _ := os.Hostname()
	cpu := runtime.NumCPU()
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			userAgent := r.Header.Get("User-Agent")
			log.Printf("%s - [%s %s %s] 200 [%s]",
				r.RemoteAddr, r.Method, r.URL.Path, r.Proto, userAgent)
			w.Header().Set("Server", "Hello/v2")

			fmt.Fprintf(w, "CPUs: %d\n", cpu)
			fmt.Fprintf(w, "Hostname: %s\n", hostname)

			fmt.Fprintf(w, "Config file:\n")
			fmt.Fprintf(w, string(content))
		})

	log.Println("Starting Hello service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
