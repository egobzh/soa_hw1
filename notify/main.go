package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("notify service listening on :30000")
	if err := http.ListenAndServe(":30000", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
