package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		}),
	}
	server.TLSConfig = &tls.Config{MinVersion: }
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("server crashed :: %s", err.Error())
	}

}
