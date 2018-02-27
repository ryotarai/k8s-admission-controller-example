package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"k8s.io/api/admission/v1beta1"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		log.Printf("%+v", r)
		log.Printf("request body: %s", body)

		response := v1beta1.AdmissionReview{
			Response: &v1beta1.AdmissionResponse{
				Allowed: true,
			},
		}

		resp, err := json.Marshal(response)
		w.Write(resp)
	})

	listen := ":8080"
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	log.Printf("cert: %s, key: %s, listen: %s", certFile, keyFile, listen)

	log.Fatal(http.ListenAndServeTLS(listen, certFile, keyFile, nil))
}
