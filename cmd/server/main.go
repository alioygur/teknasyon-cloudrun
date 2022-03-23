package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.HandleFunc("/long-running-job", func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(15 * time.Second)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok!"))

	})

	router.HandleFunc("/dump-req", func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, false)

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "%s\n", b)

	})

	// start the server!
	log.Printf("Listening on port %s", getDefaultPort())
	if err := http.ListenAndServe(":"+getDefaultPort(), router); err != nil {
		log.Fatal(err)
	}
}

func getDefaultPort() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}