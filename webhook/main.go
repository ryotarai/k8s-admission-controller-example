package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v", r)
	})

	listen := ":8080"
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	log.Printf("cert: %s, key: %s, listen: %s", certFile, keyFile, listen)

	log.Fatal(http.ListenAndServeTLS(listen, certFile, keyFile, nil))
}
